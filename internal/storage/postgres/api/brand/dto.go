// Package brand TODO
package brand

import "github.com/sorohimm/uacs-store-back/pkg/api"

func NewBrandFromRequest(req *api.CreateBrandRequest) *Brand {
	return &Brand{
		Name: req.Name,
	}
}

type Brand struct {
	ID   int64
	Name string
}

func (o Brand) ToAPIResponse() *api.CreateBrandResponse {
	return &api.CreateBrandResponse{
		ID:   o.ID,
		Name: o.Name,
	}
}
