// Package handler TODO
package handler

import (
	"context"
	"errors"
	"time"

	jwt2 "github.com/sorohimm/uacs-store-back/internal/jwt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/sorohimm/uacs-store-back/internal/log"
	"github.com/sorohimm/uacs-store-back/internal/security"
	"github.com/sorohimm/uacs-store-back/internal/storage"
	repo "github.com/sorohimm/uacs-store-back/internal/storage/postgres/auth"
	proto "github.com/sorohimm/uacs-store-back/pkg/auth"
)

func NewAuthHandler(schema string, pool *pgxpool.Pool) *AuthHandler {
	return &AuthHandler{
		authRepoCommander: repo.NewAuthRepo(schema, pool),
		authRepoRequester: repo.NewAuthRepo(schema, pool),
	}
}

type AuthHandler struct {
	proto.UnimplementedAuthServiceServer
	authRepoCommander     storage.UserCommander
	authRepoRequester     storage.UserRequester
	accessExpireDuration  time.Duration
	refreshExpireDuration time.Duration
	signingKey            string
}

func (o *AuthHandler) SetSigningKey(signingKey string) *AuthHandler {
	o.signingKey = signingKey
	return o
}

func (o *AuthHandler) SetAccessExpireDuration(expireDuration time.Duration) *AuthHandler {
	o.accessExpireDuration = expireDuration
	return o
}

func (o *AuthHandler) SetRefreshExpireDuration(expireDuration time.Duration) *AuthHandler {
	o.refreshExpireDuration = expireDuration
	return o
}

func (o *AuthHandler) Registration(ctx context.Context, req *proto.RegistrationRequest) (*empty.Empty, error) {
	logger := log.FromContext(ctx).Sugar()
	logger.Debug("AuthHandler.Registration was called")

	var (
		err error
		_   *repo.User
	)

	salt := security.GenerateSalt(req.Password)
	saltedPwd := security.SaltPassword(req.Password, salt)
	hashPwd, err := security.HashPassword(saltedPwd)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	repoReq := repo.CreateUserRequest{
		User: repo.User{
			Username: req.Username,
			Email:    req.Email,
			Password: hashPwd,
			Role:     req.Role,
		},
		PwdSalt: salt,
	}

	if _, err = o.authRepoCommander.CreateUser(ctx, &repoReq); err != nil {
		if errors.Is(err, repo.ErrUserAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (o *AuthHandler) Login(ctx context.Context, req *proto.LoginRequest) (*empty.Empty, error) {
	logger := log.FromContext(ctx).Sugar()
	logger.Debug("AuthHandler.Login was called")

	var (
		credentials *repo.Credentials
		err         error
	)
	if credentials, err = o.authRepoRequester.GetUserCredentialByUsername(ctx, req.Username); err != nil {
		if errors.Is(err, repo.ErrNotFound) {
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	saltedReqPwd := security.SaltPassword(req.Password, credentials.PwdSalt)

	if ok := security.DoPasswordsMatch(credentials.Password, saltedReqPwd); !ok {
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}

	pair, err := jwt2.GenerateTokenPair(o.accessExpireDuration, o.refreshExpireDuration, o.signingKey, credentials.UserID, "")
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = jwt2.SetAccessTokenInContext(ctx, pair.AccessToken); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if err = jwt2.SetRefreshTokenInContext(ctx, pair.RefreshToken); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (o *AuthHandler) Logout(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return &empty.Empty{}, nil
}

func (o *AuthHandler) RefreshAccessToken(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	rt, err := jwt2.GetRefreshTokenFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if ok := jwt2.IsValidToken(rt, []byte(o.signingKey)); !ok {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	t, err := jwt2.GetAccessTokenFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	claims, err := jwt2.ParseToken(t, []byte(o.signingKey))
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	newT, err := jwt2.GenerateAccessToken(o.accessExpireDuration, o.signingKey, claims.UserID, claims.UserRole)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = jwt2.SetAccessTokenInContext(ctx, newT); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = jwt2.SetRefreshTokenInContext(ctx, rt); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}
