package storage

import (
	"context"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/brand"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/category"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/product/dto"
	auth "github.com/sorohimm/uacs-store-back/internal/storage/postgres/auth"
	"github.com/sorohimm/uacs-store-back/pkg/api"
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

type BrandCommander interface {
	CreateBrand(ctx context.Context, request *api.CreateBrandRequest) (*brand.Brand, error)
}

type CategoryCommander interface {
	CreateCategory(ctx context.Context, request *api.CreateCategoryRequest) (*category.Category, error)
}

type UserCommander interface {
	CreateUser(ctx context.Context, req *auth.CreateUserRequest) (*auth.User, error)
}

type UserRequester interface {
	GetUserByID(ctx context.Context, userID int64) (*auth.User, error)
	GetUserByUsername(ctx context.Context, username string) (*auth.User, error)
	GetUserCredentialByUsername(ctx context.Context, username string) (*auth.Credentials, error)
}
