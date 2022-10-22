package product

import (
	"context"

	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/product/dto"

	"github.com/sorohimm/uacs-store-back/pkg/product"
)

//go:generate mockgen -source=product-handler.go -package=model -destination=product-handler_mock.go

type ProductRequesterHandler interface {
	GetProductByID(ctx context.Context, id int64) (*dto.Product, error)
	GetAllProducts(ctx context.Context, limit int64, offset int64) (*dto.Products, error)
	GetAllProductsWithBrand(ctx context.Context, brandID int64, limit int64, offset int64) (*dto.Products, error)
	GetAllProductsWithType(ctx context.Context, typeID int64, limit int64, offset int64) (*dto.Products, error)
	GetAllProductsWithBrandAndType(ctx context.Context, typeID int64, brandID int64, limit int64, offset int64) (*dto.Products, error)
}

type ProductCommanderHandler interface {
	CreateProduct(ctx context.Context, req *product.CreateProductRequest) (*dto.Product, error)
}
