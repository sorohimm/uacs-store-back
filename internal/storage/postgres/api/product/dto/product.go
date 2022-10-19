package dto

import "github.com/sorohimm/uacs-store-back/pkg/api"

func NewProductFromRequest(req *api.CreateProductRequest) *Product {
	var info []*ProductInfo
	for _, el := range req.Info {
		info = append(info, NewProductInfoFromAPI(el))
	}

	return &Product{
		Name:    req.Name,
		Price:   req.Price,
		BrandID: req.BrandId,
		TypeID:  req.TypeId,
		Info:    info,
	}
}

type Product struct {
	ID      int64
	Name    string
	Price   float32
	BrandID int64
	TypeID  int64
	Img     string
	Info    []*ProductInfo
}

func (o *Product) SetID(id int64) *Product {
	o.ID = id
	return o
}

func (o Product) ToAPIResponse() *api.ProductResponse {
	var info []*api.ProductInfo
	for _, el := range o.Info {
		info = append(info, el.ToAPI())
	}
	return &api.ProductResponse{
		Id:    o.ID,
		Name:  o.Name,
		Price: o.Price,
		Img:   o.Img,
		Info:  info,
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
