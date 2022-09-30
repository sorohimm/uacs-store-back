package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

var (
	AccessTokenKey  = "access_token"
	RefreshTokenKey = "refresh_token"
)

// ParseToken returns user's username from jwt token
func ParseToken(accessToken string, signingKey []byte) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.ID, nil
	}

	return "", ErrInvalidAccessToken
}

func GenerateTokenPair(accessExpireDuration, refreshExpireDuration time.Duration, secret string, id int64, role string) (map[string]string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = 1
	claims["id"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(accessExpireDuration).Unix()

	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(refreshExpireDuration).Unix()

	rt, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		AccessTokenKey:  t,
		RefreshTokenKey: rt,
	}, nil
}
