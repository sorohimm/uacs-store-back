package handler

import (
	"context"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/sorohimm/uacs-store-back/internal/service/rbac/auth"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/emptypb"

	rbacRepo "github.com/sorohimm/uacs-store-back/internal/storage/postgres/rbac"
	rbac "github.com/sorohimm/uacs-store-back/pkg/rbac"
)

func NewAuthHandler(schema string, pool *pgxpool.Pool) *AuthHandler {
	return &AuthHandler{}
}

type AuthHandler struct {
	rbac.UnimplementedAuthServiceServer
	authRepo       rbacRepo.AuthRepo
	expireDuration time.Duration
}

func (o *AuthHandler) Registration(ctx context.Context, req *rbac.RegistrationRequest) (*empty.Empty, error) {
	return nil, nil
}

func (o *AuthHandler) Login(ctx context.Context, req *rbac.LoginRequest) (*empty.Empty, error) {
	return nil, nil
}

func (o *AuthHandler) Logout(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, nil
}

func (o *AuthHandler) token(username string) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &auth.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(o.expireDuration)),
			IssuedAt:  jwt.At(time.Now()),
		},
		Username: username,
	})

	return token
}
