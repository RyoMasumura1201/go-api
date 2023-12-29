package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setUp(){
	mux = http.NewServeMux()
	mux.HandleFunc("/todo/", handleTodo)
	mux.HandleFunc("/todos/", handleTodos)
	writer = httptest.NewRecorder()
}

func tearDown(){
}

func TestHandleGetTodos(t *testing.T){
	request, _ := http.NewRequest("GET", "/todos/", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestHandlePostTodo(t *testing.T){
	json := strings.NewReader(`{"content":"buy coffee beans"}`)
	request, _:= http.NewRequest("POST", "/post/", json)
	mux.ServeHTTP(writer, request)

	if writer.Code !=200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}