package request

import "errors"

const (
	loginPath    = "/auth/login"
	registerPath = "/auth/registration"
)

var ErrBadStatusCode = errors.New("bad status code")
