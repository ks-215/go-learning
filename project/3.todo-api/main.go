package main

import (
	"todo-api_lesson/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/todos", handler.CreateTodo)
	r.GET("/todos/:id", handler.GetTodo)
	r.GET("/todos", handler.GetTodos)
	r.DELETE("/todos/:id", handler.DeleteTodo)
	r.PUT("/todos/:id", handler.CompleteTodo)

	r.Run(":7777")
}
