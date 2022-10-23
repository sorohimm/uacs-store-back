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

func NewCartCommanderHandler(schema string, pool *pgxpool.Pool) *CartCommanderHandler {
	return &CartCommanderHandler{
		commander: cart.NewRepo(schema, pool),
		requester: cart.NewRepo(schema, pool),
	}
}

type CartCommanderHandler struct {
	api.UnimplementedCartServiceServer
	commander storage.CartCommander
	requester storage.CartRequester
}

func (o *CartCommanderHandler) AddCartItem(ctx context.Context, req *api.CartItem) (*emptypb.Empty, error) {
	if _, err := o.commander.AddCartItem(ctx, req); err != nil {
		if errors.Is(err, postgres.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return nil, nil
}

func (o *CartCommanderHandler) DeleteCartItem(ctx context.Context, req *api.CartItem) (*emptypb.Empty, error) {
	if err := o.commander.DeleteCartItem(ctx, req); err != nil {
		if errors.Is(err, postgres.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return nil, nil
}

func (o *CartCommanderHandler) PatchCartItem(ctx context.Context, req *api.CartItem) (*api.Cart, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchCartItem not implemented")
}

func (o *CartCommanderHandler) GetCart(ctx context.Context, req *api.CartReq) (*api.Cart, error) {
	if _, err := o.requester.GetCart(ctx, req); err != nil {
		if errors.Is(err, postgres.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return nil, nil
}
