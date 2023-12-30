package main

type FakeTodo struct {
	Id int
	Content string
}

func (todo *FakeTodo)fetch(id int)(err error){
	todo.Id = id
	return
}

func (todo *FakeTodo)create()(err error){
	return
}

func (todo *FakeTodo)update()(err error){
	return
}

func (todo *FakeTodo)delete()(err error){
	return
}