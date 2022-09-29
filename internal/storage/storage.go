package storage

import (
	"context"
	"github.com/sorohimm/shop/internal/storage/postgres/api/product/dto"

	"github.com/sorohimm/shop/pkg/api"
)

type ProductRequester interface {
	GetProductByID(ctx context.Context, id int64) (*dto.Product, error)
	GetAllProducts(ctx context.Context, limit int64, offset int64) (*dto.Products, error)
	GetAllProductsWithBrand(ctx context.Context, brandID int64, limit int64, offset int64) (*dto.Products, error)
	GetAllProductsWithType(ctx context.Context, typeID int64, limit int64, offset int64) (*dto.Products, error)
	GetAllProductsWithBrandAndType(ctx context.Context, typeID int64, brandID int64, limit int64, offset int64) (*dto.Products, error)
}

type ProductCommander interface {
	CreateProduct(ctx context.Context, request *api.CreateProductRequest) (*dto.Product, error)
}

type InfoRequester interface {
	GetInfo(ctx context.Context, productID string) (*dto.ProductInfo, error)
}

type InfoCommander interface {
	AddInfo(ctx context.Context, info []*api.ProductInfo, productID int64) error
}
