package user

import (
	"time"

	"github.com/korroziea/photo-storage/internal/domain"
)

type signUpReq struct {
	FirstName string `json:"first_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func (r *signUpReq) toDomain() domain.User {
	u := domain.User{
		FirstName: r.FirstName,
		Email:     r.Email,
		Password:  r.Password,
	}

	return u
}

type signInReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *signInReq) toDomain() domain.User {
	u := domain.User{
		Email:    r.Email,
		Password: r.Password,
	}

	return u
}

type userResp struct {
	FirstName string    `json:"first_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at" binding:"omitempty"`
}
