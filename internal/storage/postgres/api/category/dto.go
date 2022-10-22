package category

import "github.com/sorohimm/uacs-store-back/pkg/product"

func NewCategoryFromRequest(req *product.CreateCategoryRequest) *Category {
	return &Category{
		Name: req.GetName(),
	}
}

type Category struct {
	ID   int64
	Name string
}

func (o Category) ToAPIResponse() *product.CreateCategoryResponse {
	return &product.CreateCategoryResponse{
		Id:   o.ID,
		Name: o.Name,
	}
}
