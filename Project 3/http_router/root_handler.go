package http_router

import (
	"net/http"
	
	"github.com/gin-gonic/gin"

	"gin-example/model"
)


func root_handler(c *gin.Context) {
	c.JSON(http.StatusOK, model.Router{
		Name: "gin",
	})
}