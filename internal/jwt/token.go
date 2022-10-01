package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

// ParseToken returns user's id from jwt token
func ParseToken(tokenString string, signingKey []byte) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidAccessToken
}

func IsValidToken(tokenString string, signingKey []byte) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})
	if err == nil && token.Valid {
		return true
	}

	return false
}

func GenerateTokenPair(accessExpireDuration, refreshExpireDuration time.Duration, secret string, id int64, role string) (*TokenPair, error) {

	token, err := GenerateAccessToken(accessExpireDuration, secret, id, role)
	if err != nil {
		return nil, err
	}
	rt, err := GenerateRefreshToken(refreshExpireDuration, secret)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  token,
		RefreshToken: rt,
	}, nil
}

func GenerateAccessToken(expireDuration time.Duration, secret string, id int64, role string) (string, error) {
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(expireDuration)),
			IssuedAt:  jwt.At(time.Now()),
		},
		UserID:   id,
		UserRole: role,
	})

	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func GenerateRefreshToken(expireDuration time.Duration, secret string) (string, error) {
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(expireDuration)),
			IssuedAt:  jwt.At(time.Now()),
		},
	)

	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return rt, nil
}
