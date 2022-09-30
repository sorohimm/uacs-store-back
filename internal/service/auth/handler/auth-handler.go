package handler

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/sorohimm/uacs-store-back/internal/log"
	jwt2 "github.com/sorohimm/uacs-store-back/internal/service/auth/jwt"
	"github.com/sorohimm/uacs-store-back/internal/storage"
	rbacRepo "github.com/sorohimm/uacs-store-back/internal/storage/postgres/auth"
	rbac "github.com/sorohimm/uacs-store-back/pkg/auth"
)

func NewAuthHandler(schema string, pool *pgxpool.Pool) *AuthHandler {
	return &AuthHandler{
		authRepo: rbacRepo.NewAuthRepo(schema, pool, ""),
	}
}

type AuthHandler struct {
	rbac.UnimplementedAuthServiceServer
	authRepo       storage.AuthCommander
	expireDuration time.Duration
}

func (o *AuthHandler) Registration(ctx context.Context, req *rbac.RegistrationRequest) (*empty.Empty, error) {
	logger := log.FromContext(ctx).Sugar()
	logger.Debug("AuthHandler.Registration was called")

	var (
		err error
		_   *rbacRepo.User
	)
	if _, err = o.authRepo.Registration(ctx, req); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &empty.Empty{}, status.Error(codes.OK, "success")
}

func (o *AuthHandler) Login(ctx context.Context, req *rbac.LoginRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (o *AuthHandler) Logout(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &empty.Empty{}, nil
}

func (o *AuthHandler) getToken(username string) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt2.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(o.expireDuration)),
			IssuedAt:  jwt.At(time.Now()),
		},
		Username: username,
	})

	return token
}

func (o *AuthHandler) IsAuthorized(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &empty.Empty{}, nil
}
