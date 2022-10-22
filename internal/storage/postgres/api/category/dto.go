package category

func NewCategoryFromRequest(req *product.CreateCategoryRequest) *Category {
	return &Category{
		Name: req.GetName(),
	}
}

type Category struct {
	ID   int64
	Name string
}

func (o Category) ToAPIResponse() *product.CategoryResponse {
	return &product.CategoryResponse{
		Id:   o.ID,
		Name: o.Name,
	}
}
