package product

import "github.com/sorohimm/shop/pkg/api"

type Product struct {
	Id    int64
	Name  string
	Price float32
	Img   string
}

func (o *Product) ToAPIResponse() *api.ProductResponse {
	return &api.ProductResponse{
		Id:    o.Id,
		Name:  o.Name,
		Price: o.Price,
		Img:   o.Img,
	}
}

type Products []*Product

func (o Products) ToAPIResponse() *api.AllProductsResponse {
	var products []*api.ProductResponse
	for _, el := range o {
		products = append(products, el.ToAPIResponse())
	}

	return &api.AllProductsResponse{Products: products}
}
