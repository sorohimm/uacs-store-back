// Package jwt TODO
package jwt

import "time"

type Config struct {
	Secret                     string        `long:"secret" env:"SECRET" description:"jwt secret encryption key"`
	AccessTokenExpireDuration  time.Duration `long:"access-duration" env:"ACCESS_EXPIRE_DURATION" default:"3h" description:"access token expire duration"`
	RefreshTokenExpireDuration time.Duration `long:"refresh-duration" env:"REFRESH_EXPIRE_DURATION" default:"168h" description:"refresh token expire duration"`
}
