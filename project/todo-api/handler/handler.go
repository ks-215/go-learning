package handler

import (
	"strconv"
	"sync"
	"todo-api_lesson/model"

	"github.com/gin-gonic/gin"
)

var (
	todos  []model.Task
	nextID int = 1
	mu     sync.Mutex
)

// POST
func CreateTodo(c *gin.Context) {

	mu.Lock()
	defer mu.Unlock()

	var req model.Task

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "不正リクエスト"})
		return
	}

	req.ID = nextID
	req.Completed = false
	nextID++
	todos = append(todos, req)

	c.JSON(201, req)
}

// GET
func GetTodo(c *gin.Context) {

	mu.Lock()
	defer mu.Unlock()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"message": "不正リクエスト"})
		return
	}

	for _, todo := range todos {
		if todo.ID == id {
			c.JSON(200, todo)
			return
		}
	}

	c.JSON(404, gin.H{"error": "見つかりませんでした"})
}

// GET(all)
func GetTodos(c *gin.Context) {

	mu.Lock()
	defer mu.Unlock()

	c.JSON(200, todos)
}

// DELETE
func DeleteTodo(c *gin.Context) {

	mu.Lock()
	defer mu.Unlock()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"message": "不正リクエスト"})
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(200, gin.H{"message": "削除しました"})
			return
		}
	}

	c.JSON(404, gin.H{"message": "見つかりませんでした"})
}

// PUT
func CompleteTodo(c *gin.Context) {

	mu.Lock()
	defer mu.Unlock()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"message": "不正リクエスト"})
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Completed = true
			c.JSON(200, todos[i])
			return
		}
	}
	c.JSON(404, gin.H{"message": "見つかりませんでした"})
}
