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
	"github.com/sorohimm/uacs-store-back/internal/security"
	jwt2 "github.com/sorohimm/uacs-store-back/internal/service/auth/jwt"
	"github.com/sorohimm/uacs-store-back/internal/storage"
	repo "github.com/sorohimm/uacs-store-back/internal/storage/postgres/auth"
	protoAuth "github.com/sorohimm/uacs-store-back/pkg/auth"
)

func NewAuthHandler(schema string, pool *pgxpool.Pool) *AuthHandler {
	return &AuthHandler{
		authRepoCommander: repo.NewAuthRepo(schema, pool, ""),
		authRepoRequester: repo.NewAuthRepo(schema, pool, ""),
	}
}

type AuthHandler struct {
	protoAuth.UnimplementedAuthServiceServer
	authRepoCommander storage.UserCommander
	authRepoRequester storage.UserRequester
	expireDuration    time.Duration
	signingKey        []byte
}

func (o *AuthHandler) SetSigningKey(signingKey []byte) *AuthHandler {
	o.signingKey = signingKey
	return o
}

func (o *AuthHandler) SetExpireDuration(expireDuration time.Duration) *AuthHandler {
	o.expireDuration = expireDuration
	return o
}

func (o *AuthHandler) Registration(ctx context.Context, req *protoAuth.RegistrationRequest) (*empty.Empty, error) {
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

func (o *AuthHandler) Login(ctx context.Context, req *protoAuth.LoginRequest) (*empty.Empty, error) {
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
		return nil, err
	}

	if !security.DoPasswordsMatch(hashReqPwd, credentials.Password) {
		return nil, status.Errorf(codes.PermissionDenied, err.Error())
	}

	token := o.genToken(credentials.UserID)
	stringToken, _ := token.SignedString(o.signingKey)

	if err = SetTokenInContext(ctx, stringToken); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (o *AuthHandler) Logout(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &empty.Empty{}, nil
}

func (o *AuthHandler) IsAuthorized(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &empty.Empty{}, nil
}

func (o *AuthHandler) genToken(id int64) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt2.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(o.expireDuration)),
			IssuedAt:  jwt.At(time.Now()),
		},
		UserID: id,
	})

	return token
}
