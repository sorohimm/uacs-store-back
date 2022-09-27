package storage

import (
	"context"
	"github.com/sorohimm/shop/pkg/api"

	"github.com/sorohimm/shop/internal/storage/postgres/api/product"
)

type ProductRequester interface {
	GetProductById(ctx context.Context, id int64) (*product.Product, error)
	GetAllProducts(ctx context.Context) (product.Products, error)
}

type ProductCommander interface {
	CreateProduct(ctx context.Context, request *api.CreateProductRequest) (*product.Product, error)
}
