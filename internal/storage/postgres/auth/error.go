package auth

import "errors"

var ErrNotFound = errors.New("not found")
var ErrUserAlreadyExists = errors.New("user already exists")
