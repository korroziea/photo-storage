package user

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/korroziea/photo-storage/internal/domain"
	"github.com/korroziea/photo-storage/internal/handler/response"
	"go.uber.org/zap"
)

type Service interface {
	SignUp(ctx context.Context, user domain.User) (domain.User, error)
	SignIn(ctx context.Context, user domain.User) (domain.User, error)
}

type Handler struct {
	l       zap.Logger
	service Service
}

func New(l zap.Logger, service Service) *Handler {
	h := &Handler{
		l:       l,
		service: service,
	}

	return h
}

func (h *Handler) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var signUpReq signUpReq
		if err := c.ShouldBindJSON(&signUpReq); err != nil {
			h.l.Error("ShouldBindJSON", zap.Error(err))
			
			response.Error(c, err)
			
			return
		}

		resp, err := h.service.SignUp(c.Request.Context(), signUpReq.toDomain())
		if err != nil {
			h.l.Error("service.SignUp", zap.Error(err))
			
			response.Error(c, err)

			return
		}

		c.JSON(http.StatusCreated, userResp{
			FirstName: resp.FirstName,
			Email:     resp.Email,
			CreatedAt: resp.CreatedAt,
		})
	}
}

func (h *Handler) SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var signInReq signInReq
		if err := c.ShouldBindJSON(&signInReq); err != nil {
			response.Error(c, err)

			return
		}

		resp, err := h.service.SignIn(c.Request.Context(), signInReq.toDomain())
		if err != nil {
			response.Error(c, err)

			return
		}

		c.JSON(http.StatusOK, userResp{
			FirstName: resp.FirstName,
			Email:     resp.Email,
		})
	}
}
