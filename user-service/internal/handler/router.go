package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/korroziea/photo-storage/internal/handler/user"
)

type Handler struct {
	user *user.Handler
}

func New(user *user.Handler) *Handler {
	h := &Handler{
		user: user,
	}
	
	return h
}

func (h *Handler) InitRoutes() http.Handler {
	r := gin.New()
	
	r.POST("/ping", ping())
	
	r.POST("/sign-up", h.user.SignUp())
	r.POST("/sign-in", h.user.SignIn())
	
	return r
}

func ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Pong")
	}
}
