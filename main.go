package main

import (
	"godo/controllers"
	"godo/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.ConnectDataBase()

	r.GET("/", controllers.GetAllTodo)
	r.POST("/", controllers.CerateTodo)

	r.Run(":8000")
}
