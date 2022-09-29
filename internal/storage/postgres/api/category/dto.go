package category

import "github.com/sorohimm/uacs-store-back/pkg/api"

type Category struct {
	ID   int64
	Name string
}

func (o Category) ToAPIResponse() *api.CreateCategoryResponse {
	return &api.CreateCategoryResponse{
		Id:   o.ID,
		Name: o.Name,
	}
}
