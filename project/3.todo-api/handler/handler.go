package handler

import (
	"context"
	"strconv"
	"todo-api_lesson/db"
	"todo-api_lesson/model"

	"github.com/gin-gonic/gin"
)

// POST
func CreateTodo(c *gin.Context) {
	var req model.Task

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "不正リクエスト"})
		return
	}

	err := db.DB.QueryRow(
		context.Background(),
		"INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id",
		req.Title, false,
	).Scan(&req.ID)

	if err != nil {
		c.JSON(500, gin.H{"error": "DBエラー"})
		return
	}

	req.Completed = false
	c.JSON(201, req)
}

// GET
func GetTodo(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"message": "不正リクエスト"})
		return
	}

	var todo model.Task
	err = db.DB.QueryRow(
		context.Background(),
		"SELECT id, title, completed FROM todos WHERE id = $1",
		id,
	).Scan(&todo.ID, &todo.Title, &todo.Completed)

	if err != nil {
		c.JSON(404, gin.H{"error": "見つかりませんでした"})
		return
	}

	c.JSON(200, todo)
}

// GET(all)
func GetTodos(c *gin.Context) {
	rows, err := db.DB.Query(
		context.Background(),
		"SELECT id, title, completed FROM todos",
	)

	if err != nil {
		c.JSON(500, gin.H{"error": "DBエラー"})
		return
	}
	defer rows.Close()

	var todos []model.Task
	for rows.Next() {
		var todo model.Task
		if err = rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			c.JSON(500, gin.H{"error": "DBエラー"})
			return
		}
		todos = append(todos, todo)
	}

	c.JSON(200, todos)
}

// DELETE
func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"message": "不正リクエスト"})
		return
	}

	result, err := db.DB.Exec(
		context.Background(),
		"DELETE FROM todos WHERE id = $1",
		id,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": "DBエラー"})
		return
	}

	if result.RowsAffected() == 0 {
		c.JSON(404, gin.H{"message": "見つかりませんでした"})
		return
	}

	c.JSON(200, gin.H{"message": "削除しました"})
}

// PUT
func CompleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"message": "不正リクエスト"})
		return
	}

	var todo model.Task
	err = db.DB.QueryRow(
		context.Background(),
		"UPDATE todos SET completed = true WHERE id = $1 RETURNING id, title, completed",
		id,
	).Scan(&todo.ID, &todo.Title, &todo.Completed)

	if err != nil {
		c.JSON(404, gin.H{"message": "見つかりませんでした"})
		return
	}

	c.JSON(200, todo)
}
