package handler

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/sorohimm/shop/internal/storage"
	"github.com/sorohimm/shop/internal/storage/postgres/api/product"
	"github.com/sorohimm/shop/pkg/api"
)

func NewStoreRequesterHandler(schema string, pool *pgxpool.Pool) *StoreRequesterHandler {
	return &StoreRequesterHandler{
		productRequester: product.NewProductRepo(schema, pool),
	}
}

type StoreRequesterHandler struct {
	api.UnimplementedStoreServiceRequesterServer
	productRequester storage.ProductRequester
}

func (o *StoreRequesterHandler) GetProduct(ctx context.Context, req *api.ProductRequest) (*api.ProductResponse, error) {
	prod, err := o.productRequester.GetProductById(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, product.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return prod.ToAPIResponse(), nil
}

func (o *StoreRequesterHandler) GetAllProducts(ctx context.Context, req *api.AllProductsRequest) (*api.AllProductsResponse, error) {
	return nil, nil
}
