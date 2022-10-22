// Package brand TODO
package brand

func NewBrandFromRequest(req *product.CreateBrandRequest) *Brand {
	return &Brand{
		Name: req.Name,
	}
}

type Brand struct {
	ID   int64
	Name string
}

func (o Brand) ToAPIResponse() *product.BrandResponse {
	return &product.BrandResponse{
		Id:   o.ID,
		Name: o.Name,
	}
}
