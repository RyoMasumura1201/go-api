package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Text interface {
	fetch(id int) (err error)
	create() (err error)
	update() (err error)
	delete() (err error)
	fetchAll() (todos []Todo, err error)
}

type Todo struct {
	Db *sql.DB `json:"-"`
	Id int `json:"id"`
	Content string `json:"content"`
}

func (todo *Todo)fetch(id int)(err error){
	err = todo.Db.QueryRow("select id, content from todos where id = $1", id).Scan(&todo.Id, &todo.Content)
	return
}

func (todo *Todo)fetchAll()(todos []Todo, err error){
	rows, err := todo.Db.Query("select id, content from todos")
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
	stmt, err := todo.Db.Prepare(statement)
	if err !=nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(todo.Content).Scan(&todo.Id)
	return
}

func (todo *Todo)update()(err error){
	_, err = todo.Db.Exec("update todos set content = $2 where id = $1", todo.Id, todo.Content)
	return
}

func (todo *Todo)delete()(err error){
	_, err = todo.Db.Exec("delete from todos where id = $1", todo.Id)
	return
}
