package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	hub := newHub()
	go hub.run()

	router := gin.Default()
	router.LoadHTMLGlob("static/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/ws", func(c *gin.Context) {
		serveWs(hub, c)
	})

	router.Run(":8080")
}
