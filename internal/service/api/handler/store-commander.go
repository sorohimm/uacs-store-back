// Package handler TODO
package handler

import (
	"context"

	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/brand"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/category"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/product/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/sorohimm/uacs-store-back/internal/storage"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/product"
	"github.com/sorohimm/uacs-store-back/pkg/api"
)

func NewStoreCommanderHandler(schema string, pool *pgxpool.Pool) *StoreCommanderHandler {
	return &StoreCommanderHandler{
		productCommander:  product.NewProductRepo(schema, pool),
		brandCommander:    brand.NewBrandRepo(schema, pool),
		categoryCommander: category.NewCategoryRepo(schema, pool),
	}
}

type StoreCommanderHandler struct {
	api.UnimplementedStoreServiceCommanderServer
	productCommander  storage.ProductCommander
	brandCommander    storage.BrandCommander
	categoryCommander storage.CategoryCommander
}

func (o *StoreCommanderHandler) CreateProduct(ctx context.Context, req *api.CreateProductRequest) (*api.ProductResponse, error) {
	var (
		prod *dto.Product
		err  error
	)

	if prod, err = o.productCommander.CreateProduct(ctx, req); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return prod.ToAPIResponse(), nil
}

func (o *StoreCommanderHandler) CreateCategory(ctx context.Context, req *api.CreateCategoryRequest) (*api.CreateCategoryResponse, error) {
	var (
		newCategory *category.Category
		err         error
	)

	if newCategory, err = o.categoryCommander.CreateCategory(ctx, req); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return newCategory.ToAPIResponse(), nil
}

func (o *StoreCommanderHandler) CreateBrand(ctx context.Context, req *api.CreateBrandRequest) (*api.CreateBrandResponse, error) {
	var (
		newBrand *brand.Brand
		err      error
	)

	if newBrand, err = o.brandCommander.CreateBrand(ctx, req); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return newBrand.ToAPIResponse(), nil
}
