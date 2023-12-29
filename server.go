package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

type Todo struct {
	Id int
	Content string
}


func handleRequest(w http.ResponseWriter, r *http.Request){
	var err error
	switch r.Method {
	case "POST":
		err = handlePost(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func handlePost(w http.ResponseWriter, r *http.Request)(err error){
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

func handleDelete(w http.ResponseWriter, r *http.Request)(err error){
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

func main(){
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/todo/", handleRequest)
	server.ListenAndServe()
}