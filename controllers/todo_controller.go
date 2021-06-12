package controllers

import (
	"godo/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": services.GetAllTodo()})
}

func CerateTodo(c *gin.Context) {
	var input services.CreateTodoDto

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": services.CreateTodo(&input)})
}
