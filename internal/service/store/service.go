package store

import (
	"context"
	"github.com/sorohimm/shop/internal"
)

func NewService() *Service {
	return &Service{
		internal.NewRunGroup(),
	}
}

type Service struct {
	*internal.RunGroup
}

func (o *Service) Init(ctx context.Context, appName, version, built string) {

}
