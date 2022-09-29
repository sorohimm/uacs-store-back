package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	rbac "github.com/sorohimm/uacs-store-back/pkg/rbac"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewAuthHandler(schema string) *AuthHandler {
	return &AuthHandler{}
}

type AuthHandler struct {
	rbac.UnimplementedAuthServiceServer
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
