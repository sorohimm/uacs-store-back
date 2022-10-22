package category

import "github.com/sorohimm/uacs-store-back/pkg/api"

func NewCategoryFromRequest(req *api.CreateCategoryRequest) *Category {
	return &Category{
		Name: req.GetName(),
	}
}

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
