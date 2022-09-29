package model

import (
	"context"

	"github.com/sorohimm/uacs-store-back/pkg/api"
)

//go:generate mockgen -source=brand-handler.go -package=model -destination=brand-handler_mock.go

type CategoryRequesterHandler interface {
	CreateCategory(ctx context.Context, req *api.CreateCategoryRequest) (*api.CreateCategoryResponse, error)
}
