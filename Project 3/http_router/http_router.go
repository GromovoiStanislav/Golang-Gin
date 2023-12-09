package http_router

import (
	"strconv"
	"log"

	"github.com/gin-gonic/gin"
)

func GinNew(port int) {
	g := gin.New()

    g.GET("/", root_handler)
	
	log.Println("serving with GIN on ", port)
	go log.Fatal(g.Run(":" + strconv.Itoa(port)))
}