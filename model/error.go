package model

import "errors"

var (
	ErrNotFound     = errors.New("resource not found")
	ErrInvalidToken = errors.New("invalid jwt token")
	ErrUnauthorized = errors.New("unauthorized")
)
