package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

type Todo struct {
	Id int `json:"id"`
	Content string `json:"content"`
}


func handleTodo(w http.ResponseWriter, r *http.Request){
	var err error
	switch r.Method {
	case "POST":
		err = handlePostTodo(w, r)
	case "DELETE":
		err = handleDeleteTodo(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func handleTodos(w http.ResponseWriter, r *http.Request){
	var err error
	switch r.Method {
	case "GET":
		err = handleGetTodos(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handlePostTodo(w http.ResponseWriter, r *http.Request)(err error){
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var todo Todo
	json.Unmarshal(body, &todo)
	err = todo.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleDeleteTodo(w http.ResponseWriter, r *http.Request)(err error){
	id ,err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	todo, err := retrieve(id)
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

func handleGetTodos(w http.ResponseWriter, r *http.Request)(err error){
	todos, err := retrieveAll()
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
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/todo/", handleTodo)
	http.HandleFunc("/todos", handleTodos)
	server.ListenAndServe()
}