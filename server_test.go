package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// func TestHandleGetTodos(t *testing.T){
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/todos/", handleTodo(&FakeTodo{}))
// 	writer := httptest.NewRecorder()
// 	request, _ := http.NewRequest("GET", "/todos/", nil)
// 	mux.ServeHTTP(writer, request)

// 	if writer.Code != 200 {
// 		t.Errorf("Response code is %v", writer.Code)
// 	}
// }

func TestHandlePostTodo(t *testing.T){
	mux := http.NewServeMux()
	todo := &FakeTodo{}
	mux.HandleFunc("/todo/", handleTodo(todo))
	writer := httptest.NewRecorder()
	json := strings.NewReader(`{"content":"buy coffee beans"}`)
	request, _:= http.NewRequest("POST", "/todo/", json)
	mux.ServeHTTP(writer, request)

	if writer.Code !=200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	if todo.Content != "buy coffee beans" {
		t.Error("Content is not correct", todo.Content)
	}
}