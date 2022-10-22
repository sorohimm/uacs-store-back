// Package handler handles incoming queries
package handler

import (
	"context"
	"errors"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/brand"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/category"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/product"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/product/dto"
	"github.com/sorohimm/uacs-store-back/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/sorohimm/uacs-store-back/internal/storage"
)

func NewProductCommanderHandler(schema string, pool *pgxpool.Pool) *ProductCommanderHandler {
	return &ProductCommanderHandler{
		productCommander:  product.NewProductRepo(schema, pool),
		brandCommander:    brand.NewBrandRepo(schema, pool),
		categoryCommander: category.NewCategoryRepo(schema, pool),
	}
}

type ProductCommanderHandler struct {
	api.UnimplementedStoreServiceCommanderServer
	productCommander  storage.ProductCommander
	brandCommander    storage.BrandCommander
	categoryCommander storage.CategoryCommander
}

func (o *ProductCommanderHandler) CreateProduct(ctx context.Context, req *api.CreateProductRequest) (*api.ProductResponse, error) {
	var (
		prod *dto.Product
		err  error
	)

	if prod, err = o.productCommander.CreateProduct(ctx, req); err != nil {
		if errors.Is(err, postgres.ErrConflict) {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return prod.ToAPIResponse(), nil
}

func (o *ProductCommanderHandler) CreateCategory(ctx context.Context, req *api.CreateCategoryRequest) (*api.CategoryResponse, error) {
	var (
		newCategory *category.Category
		err         error
	)

	if newCategory, err = o.categoryCommander.CreateCategory(ctx, req); err != nil {
		if errors.Is(err, postgres.ErrConflict) {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return newCategory.ToAPIResponse(), nil
}

func (o *ProductCommanderHandler) CreateBrand(ctx context.Context, req *api.CreateBrandRequest) (*api.BrandResponse, error) {
	var (
		newBrand *brand.Brand
		err      error
	)

	if newBrand, err = o.brandCommander.CreateBrand(ctx, req); err != nil {
		if errors.Is(err, postgres.ErrConflict) {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return newBrand.ToAPIResponse(), nil
}
