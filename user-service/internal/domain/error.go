package domain

import "errors"

var (
	ErrNotFound = errors.New("entity not found")
	ErrInternal = errors.New("internal error")
	
	ErrUserAlreadyExists = errors.New("user already exists")
)
