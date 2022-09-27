package handler

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sorohimm/shop/internal/storage"
	"github.com/sorohimm/shop/internal/storage/postgres/api/product"

	"github.com/sorohimm/shop/pkg/api"
)

func NewStoreRequesterHandler(pool *pgxpool.Pool) *StoreRequesterHandler {
	return &StoreRequesterHandler{
		productRequester: product.NewProductRepo(pool),
	}
}

type StoreRequesterHandler struct {
	api.UnimplementedStoreServiceServer
	productRequester storage.ProductRequester
}

func (o *StoreRequesterHandler) GetProduct(ctx context.Context, req *api.ProductRequest) (*api.ProductResponse, error) {
	return nil, nil
}

func (o *StoreRequesterHandler) GetAllProducts(ctx context.Context, req *api.AllProductsRequest) (*api.AllProductsResponse, error) {
	return nil, nil
}
