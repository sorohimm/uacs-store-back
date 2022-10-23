// Package handler handles incoming queries
package handler

import (
	"context"
	"errors"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/product"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/product/dto"
	"github.com/sorohimm/uacs-store-back/pkg/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/sorohimm/uacs-store-back/internal/storage"
)

func NewProductRequesterHandler(schema string, pool *pgxpool.Pool) *ProductRequesterHandler {
	return &ProductRequesterHandler{
		productRequester: product.NewProductRepo(schema, pool),
	}
}

type ProductRequesterHandler struct {
	api.UnimplementedStoreServiceRequesterServer
	productRequester storage.ProductRequester
}

func (o *ProductRequesterHandler) GetProduct(ctx context.Context, req *api.ProductRequest) (*api.ProductResponse, error) {
	prod, err := o.productRequester.GetProductByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, postgres.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return prod.ToAPIResponse(), nil
}

func (o *ProductRequesterHandler) GetAllProducts(ctx context.Context, req *api.AllProductsRequest) (*api.AllProductsResponse, error) {
	limit := req.GetLimit()
	offset := offset(limit, req.GetPage())

	var (
		prod *dto.Products
		err  error
	)

	if req.GetBrandId() == 0 && req.GetTypeId() == 0 {
		prod, err = o.productRequester.GetAllProducts(ctx, limit, offset)
	}
	if req.GetBrandId() != 0 && req.GetTypeId() != 0 {
		prod, err = o.productRequester.GetAllProductsWithBrandAndType(ctx, req.GetTypeId(), req.GetBrandId(), limit, offset)
	}
	if req.GetBrandId() == 0 && req.GetTypeId() != 0 {
		prod, err = o.productRequester.GetAllProductsWithType(ctx, req.GetTypeId(), limit, offset)
	}
	if req.GetBrandId() != 0 && req.GetTypeId() == 0 {
		prod, err = o.productRequester.GetAllProductsWithBrand(ctx, req.GetBrandId(), limit, offset)
	}

	if err != nil {
		if errors.Is(err, postgres.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return prod.ToAPIResponse(), nil
}

func offset(limit int64, page int64) int64 {
	return page*limit - limit
}
