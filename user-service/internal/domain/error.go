package domain

import "errors"

var (
	ErrNotFound = errors.New("entity not found")
	ErrInternal = errors.New("internal error")

	ErrInvalidHashFormat = errors.New("invalid hash format")
	ErrInvalidHashType   = errors.New("invalid hash type")
	ErrInvalidHashVersion = errors.New("invalid hash version")

	ErrUserAlreadyExists = errors.New("user already exists")
)
