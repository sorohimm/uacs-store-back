package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/sorohimm/uacs-store-back/pkg/api"
)

type CartCommander struct {
	api.UnimplementedCartServiceCommanderServer
}

func (CartCommander) AddCartItem(ctx context.Context, req *api.CartItem) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCartItem not implemented")
}

func (CartCommander) DeleteCartItem(ctx context.Context, req *api.CartItem) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCartItem not implemented")
}
