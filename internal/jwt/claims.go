// Package jwt TODO
package jwt

import "github.com/dgrijalva/jwt-go/v4"

type Claims struct {
	jwt.StandardClaims
	UserID   int64  `json:"user_id"`
	UserRole string `json:"user_role"`
}
