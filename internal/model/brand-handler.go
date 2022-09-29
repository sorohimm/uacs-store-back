package model

import (
	"context"

	"github.com/sorohimm/shop/pkg/api"
)

//go:generate mockgen -source=brand-handler.go -package=model -destination=brand-handler_mock.go

type BrandCommanderHandler interface {
	CreateBrand(ctx context.Context, req *api.CreateBrandRequest) (*api.CreateBrandResponse, error)
}
