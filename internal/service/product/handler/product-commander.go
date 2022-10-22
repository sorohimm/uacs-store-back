// Package handler handles incoming queries
package handler

import (
	"context"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/product"
	product2 "github.com/sorohimm/uacs-store-back/pkg/api/product"

	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/brand"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/category"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/product/dto"
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
	product2.UnimplementedStoreServiceCommanderServer
	productCommander  storage.ProductCommander
	brandCommander    storage.BrandCommander
	categoryCommander storage.CategoryCommander
}

func (o *ProductCommanderHandler) CreateProduct(ctx context.Context, req *product2.CreateProductRequest) (*product2.ProductResponse, error) {
	var (
		prod *dto.Product
		err  error
	)

	if prod, err = o.productCommander.CreateProduct(ctx, req); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return prod.ToAPIResponse(), nil
}

func (o *ProductCommanderHandler) CreateCategory(ctx context.Context, req *product2.CreateCategoryRequest) (*product2.CategoryResponse, error) {
	var (
		newCategory *category.Category
		err         error
	)

	if newCategory, err = o.categoryCommander.CreateCategory(ctx, req); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return newCategory.ToAPIResponse(), nil
}

func (o *ProductCommanderHandler) CreateBrand(ctx context.Context, req *product2.CreateBrandRequest) (*product2.BrandResponse, error) {
	var (
		newBrand *brand.Brand
		err      error
	)

	if newBrand, err = o.brandCommander.CreateBrand(ctx, req); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return newBrand.ToAPIResponse(), nil
}
