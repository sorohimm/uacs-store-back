package storage

import (
	"context"

	"github.com/sorohimm/shop/internal/storage/postgres/api/product"
)

type ProductRequester interface {
	GetProductById(ctx context.Context, id int64) (*product.Product, error)
	GetAllProducts(ctx context.Context) (product.Products, error)
}
