package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init(){
	var err error
	Db, err = sql.Open("postgres", "user=gotodo dbname=gotodo password=password sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func retrieve(id int)(todo Todo, err error){
	todo = Todo{}
	err = Db.QueryRow("select id, content from todos where id = $1", id).Scan(&todo.Id, &todo.Content)
	return
}

func retrieveAll()(todos []Todo, err error){
	rows, err := Db.Query("select id, content from todos")
	if err != nil {
		return
	}
	for rows.Next(){
		todo := Todo{}
		err = rows.Scan(&todo.Id, &todo.Content)
		if err != nil {
			return
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return
}

func (todo *Todo)create()(err error){
	statement := "insert into todos (content) values ($1) returning id"
	stmt, err := Db.Prepare(statement)
	if err !=nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(todo.Content).Scan(&todo.Id)
	return
}

func (todo *Todo)update()(err error){
	_, err = Db.Exec("update todos set content = $2 where id = $1", todo.Id, todo.Content)
	return
}

func (todo *Todo)delete()(err error){
	_, err = Db.Exec("delete from todos where id = $1", todo.Id)
	return
}
