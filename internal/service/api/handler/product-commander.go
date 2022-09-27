package handler

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/sorohimm/shop/internal/storage"
	"github.com/sorohimm/shop/internal/storage/postgres/api/product"
	"github.com/sorohimm/shop/pkg/api"
)

func NewStoreCommanderHandler(schema string, pool *pgxpool.Pool) *StoreCommanderHandler {
	return &StoreCommanderHandler{
		productCommander: product.NewProductRepo(schema, pool),
	}
}

type StoreCommanderHandler struct {
	api.UnimplementedStoreServiceCommanderServer
	productCommander storage.ProductCommander
}

func (o *StoreCommanderHandler) CreateProduct(ctx context.Context, req *api.CreateProductRequest) (*api.ProductResponse, error) {
	return nil, nil
}
