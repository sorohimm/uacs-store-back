package handler

import (
	"context"
	jwt2 "github.com/sorohimm/uacs-store-back/internal/jwt"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/sessions"
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
	sessionStore          *sessions.CookieStore
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

func (o *AuthHandler) SetSessionStore(store *sessions.CookieStore) *AuthHandler {
	o.sessionStore = store
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
		return nil, err
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
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &empty.Empty{}, status.Error(codes.OK, "success")
}

func (o *AuthHandler) Login(ctx context.Context, req *proto.LoginRequest) (*empty.Empty, error) {
	var (
		credentials *repo.Credentials
		err         error
	)
	if credentials, err = o.authRepoRequester.GetUserCredentialByUsername(ctx, req.Username); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	saltedReqPwd := security.SaltPassword(req.Password, credentials.PwdSalt)
	hashReqPwd, err := security.HashPassword(saltedReqPwd)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if ok := security.DoPasswordsMatch(hashReqPwd, credentials.Password); !ok {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	pair, err := jwt2.GenerateTokenPair(o.accessExpireDuration, o.refreshExpireDuration, o.signingKey, credentials.UserID, "")
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = SetAccessTokenInContext(ctx, pair[jwt2.AccessTokenKey]); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if err = SetRefreshTokenInContext(ctx, pair[jwt2.RefreshTokenKey]); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (o *AuthHandler) Logout(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return &empty.Empty{}, nil
}

func (o *AuthHandler) RefreshToken(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return &empty.Empty{}, nil
}
