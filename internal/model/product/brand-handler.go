package product

import (
	"context"

	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/brand"

	"github.com/sorohimm/uacs-store-back/pkg/product"
)

//go:generate mockgen -source=brand-handler.go -package=model -destination=brand-handler_mock.go

type BrandCommanderHandler interface {
	CreateBrand(ctx context.Context, req *product.CreateBrandRequest) (*brand.Brand, error)
}
