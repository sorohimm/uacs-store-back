package handler

import (
	"context"
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
		productCommander: product.NewProductRepo(schema, pool),
	}
}

type StoreCommanderHandler struct {
	api.UnimplementedStoreServiceCommanderServer
	productCommander storage.ProductCommander
	infoCommander    storage.InfoCommander
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
	return nil, nil
}

func (o *StoreCommanderHandler) CreateBrand(ctx context.Context, req *api.CreateBrandRequest) (*api.CreateBrandResponse, error) {
	return nil, nil
}
