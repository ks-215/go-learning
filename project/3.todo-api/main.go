package main

import (
	"context"
	"log"
	"todo-api_lesson/db"
	"todo-api_lesson/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}
	defer db.DB.Close(context.Background())

	r := gin.Default()

	r.POST("/todos", handler.CreateTodo)
	r.GET("/todos/:id", handler.GetTodo)
	r.GET("/todos", handler.GetTodos)
	r.DELETE("/todos/:id", handler.DeleteTodo)
	r.PUT("/todos/:id", handler.CompleteTodo)

	r.Run(":8888")
}
