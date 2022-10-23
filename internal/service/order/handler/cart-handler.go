package handler

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/sorohimm/uacs-store-back/internal/storage"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/cart"
	"github.com/sorohimm/uacs-store-back/pkg/api"
)

func NewCartHandler(schema string, pool *pgxpool.Pool) *CartHandler {
	return &CartHandler{
		commander: cart.NewRepo(schema, pool),
		requester: cart.NewRepo(schema, pool),
	}
}

type CartHandler struct {
	api.UnimplementedCartServiceServer
	commander storage.CartCommander
	requester storage.CartRequester
}

func (o *CartHandler) AddCartItem(ctx context.Context, req *api.CartItem) (*emptypb.Empty, error) {
	if _, err := o.commander.AddCartItem(ctx, req); err != nil {
		if errors.Is(err, postgres.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return nil, nil
}

func (o *CartHandler) DeleteCartItem(ctx context.Context, req *api.CartItem) (*emptypb.Empty, error) {
	if err := o.commander.DeleteCartItem(ctx, req); err != nil {
		if errors.Is(err, postgres.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return nil, nil
}

func (o *CartHandler) PatchCartItem(ctx context.Context, req *api.CartItem) (*api.Cart, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchCartItem not implemented")
}

func (o *CartHandler) Info(ctx context.Context, req *api.CartReq) (*api.CartInfo, error) {
	var (
		info *cart.CartInfo
		err  error
	)
	if info, err = o.requester.GetInfo(ctx, req); err != nil {
		if errors.Is(err, postgres.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return info.ToAPI(), nil
}
