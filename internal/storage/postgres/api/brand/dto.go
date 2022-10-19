package brand

import "github.com/sorohimm/uacs-store-back/pkg/api"

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
