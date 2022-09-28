package storage

import (
	"context"

	"github.com/sorohimm/shop/pkg/api"

	"github.com/sorohimm/shop/internal/storage/postgres/api/product"
)

type ProductRequester interface {
	GetProductByID(ctx context.Context, id int64) (*product.Product, error)
	GetAllProducts(ctx context.Context, limit int64, offset int64) (*product.Products, error)
	GetAllProductsWithBrand(ctx context.Context, brandID int64, limit int64, offset int64) (*product.Products, error)
	GetAllProductsWithType(ctx context.Context, typeID int64, limit int64, offset int64) (*product.Products, error)
	GetAllProductsWithBrandAndType(ctx context.Context, typeID int64, brandID int64, limit int64, offset int64) (*product.Products, error)
}

type ProductCommander interface {
	CreateProduct(ctx context.Context, request *api.CreateProductRequest) (*product.Product, error)
}

type InfoCommander interface {
	AddInfo(ctx context.Context, info *api.ProductInfo) error
}
