package services

import "godo/db"

type CreateTodoDto struct {
	Title string `json:"title" binding:"required"`
	Done  bool   `json:"done" binding:"required"`
}

func GetAllTodo() []db.Todo {

	var todo []db.Todo

	db.DB.Find(&todo)

	return todo
}

func CreateTodo(b *CreateTodoDto) db.Todo {

	todo := db.Todo{Title: b.Title, Done: b.Done}

	db.DB.Create(&todo)

	return todo
}
