package handler

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/sorohimm/shop/pkg/api"
)

func NewStoreRequesterHandler(pool *pgxpool.Pool) *StoreRequesterHandler {
	return &StoreRequesterHandler{}
}

type StoreRequesterHandler struct {
	api.UnimplementedStoreServiceServer
}

func (o *StoreRequesterHandler) GetProduct(ctx context.Context, req *api.ProductRequest) (*api.ProductResponse, error) {
	return nil, nil
}

func (o *StoreRequesterHandler) GetAllProducts(ctx context.Context, req *api.AllProductsRequest) (*api.AllProductsResponse, error) {
	return nil, nil
}
