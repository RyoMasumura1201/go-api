package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

func handleTodo(t Text) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
	    var err error
	    switch r.Method {
	    case "POST":
	    	err = handlePostTodo(w, r, t)
	    case "DELETE":
	    	err = handleDeleteTodo(w, r, t)
	    default:
	    	http.Error(w, "Not Found", http.StatusNotFound)
	    }
	    if err != nil {
	    	http.Error(w, err.Error(), http.StatusInternalServerError)
	    	return
	    }
    }
}

func handleTodos(t Text) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
	    switch r.Method {
	    case "GET":
			err = handleGetTodos(w, r, t)
	    default:
			http.Error(w, "Not Found", http.StatusNotFound)
	    }
	    if err != nil {
	    	http.Error(w, err.Error(), http.StatusInternalServerError)
	    	return
	    }
	}
}

func handlePostTodo(w http.ResponseWriter, r *http.Request, todo Text)(err error){
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &todo)
	err = todo.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleDeleteTodo(w http.ResponseWriter, r *http.Request, todo Text)(err error){
	id ,err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	err = todo.fetch(id)
	if err != nil {
		return
	} 
	err = todo.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleGetTodos(w http.ResponseWriter, r *http.Request, todo Text)(err error){
	todos, err := todo.fetchAll()
	if err !=nil {
		return
	}
	output, err := json.MarshalIndent(&todos, "", "\t\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func main(){
	var err error
	db, err := sql.Open("postgres", "host=host.docker.internal user=gotodo dbname=gotodo password=password sslmode=disable")
	if err != nil {
		panic(err)
	}
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/todo/", handleTodo(&Todo{Db: db}))
	http.HandleFunc("/todos/", handleTodos(&Todo{Db: db}))
	server.ListenAndServe()
}