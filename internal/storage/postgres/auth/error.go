// Package auth TODO
package auth

import "errors"

var (
	ErrNotFound          = errors.New("not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)
