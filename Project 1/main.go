package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello world")
	})


	r.GET("/ping", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}