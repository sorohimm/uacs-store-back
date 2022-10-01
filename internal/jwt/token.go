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
func ParseToken(tokenString string, signingKey []byte) (int64, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, nil
	}

	return 0, ErrInvalidAccessToken
}

func IsValidToken(tokenString string, signingKey []byte) bool {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})
	if err != nil {
		return false
	}

	return true
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
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = 1
	claims["id"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(expireDuration).Unix()

	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func GenerateRefreshToken(expireDuration time.Duration, secret string) (string, error) {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(expireDuration).Unix()

	rt, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return rt, nil
}
