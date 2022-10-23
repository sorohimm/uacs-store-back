package product

import (
	"context"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/category"
	"github.com/sorohimm/uacs-store-back/pkg/api"
)

//go:generate mockgen -source=category-handler.go -package=model -destination=category-handler_mock.go

type CategoryCommanderHandler interface {
	CreateCategory(ctx context.Context, req *api.CreateCategoryRequest) (*category.Category, error)
}
