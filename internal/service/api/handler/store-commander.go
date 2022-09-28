package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	var (
		prod *product.Product
		err  error
	)

	if prod, err = o.productCommander.CreateProduct(ctx, req); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return prod.ToAPIResponse(), nil
}
