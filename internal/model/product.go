package model

import (
	"context"

	"github.com/sorohimm/shop/pkg/api"
)

//go:generate mockgen -source=product.go -package=model -destination=product_mock.go

type ProductRequesterHandler interface {
	GetProduct(ctx context.Context, req *api.ProductRequest) (*api.ProductResponse, error)
	GetAllProducts(ctx context.Context, req *api.AllProductsRequest) (*api.AllProductsResponse, error)
}

type ProductCommanderHandler interface {
	CreateProduct(ctx context.Context, req *api.CreateProductRequest) (*api.ProductResponse, error)
}
