package model

import (
	"context"

	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/product/dto"

	"github.com/sorohimm/uacs-store-back/pkg/api"
)

//go:generate mockgen -source=product-handler.go -package=model -destination=product-handler_mock.go

type ProductRequesterHandler interface {
	GetProduct(ctx context.Context, req *api.ProductRequest) (*dto.Product, error)
	GetAllProducts(ctx context.Context, req *api.AllProductsRequest) (*dto.Products, error)
}

type ProductCommanderHandler interface {
	CreateProduct(ctx context.Context, req *api.CreateProductRequest) (*dto.Product, error)
}
