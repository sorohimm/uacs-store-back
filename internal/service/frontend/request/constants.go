package request

import "errors"

const (
	loginPath             = "/auth/login"
	registerPath          = "/auth/registration"
	accessTokenHeaderKey  = "Grpc-Metadata-Access-Token"
	refreshTokenHeaderKey = "Grpc-Metadata-Refresh-Token"
)

var (
	ErrBadStatusCode     = errors.New("bad status code")
	ErrBadResponseHeader = errors.New("bad status code")
)
