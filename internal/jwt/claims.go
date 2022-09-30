package jwt

import "github.com/dgrijalva/jwt-go/v4"

type Claims struct {
	jwt.StandardClaims
	UserID int64 `json:"user_id"`
}
