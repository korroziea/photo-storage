package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes() http.Handler {
	r := gin.New()
	
	r.POST("/ping", ping())
	
	return r
}

func ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Pong")
	}
}
