package store

import "github.com/sorohimm/shop/internal"

func NewService() *Service {
	return &Service{
		internal.NewRunGroup(),
	}
}

type Service struct {
	*internal.RunGroup
}
