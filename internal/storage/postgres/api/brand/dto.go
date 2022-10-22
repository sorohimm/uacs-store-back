// Package brand TODO
package brand

import "github.com/sorohimm/uacs-store-back/pkg/product"

func NewBrandFromRequest(req *product.CreateBrandRequest) *Brand {
	return &Brand{
		Name: req.Name,
	}
}

type Brand struct {
	ID   int64
	Name string
}

func (o Brand) ToAPIResponse() *product.CreateBrandResponse {
	return &product.CreateBrandResponse{
		ID:   o.ID,
		Name: o.Name,
	}
}
