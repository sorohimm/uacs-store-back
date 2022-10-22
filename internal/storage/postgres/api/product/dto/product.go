package dto

import "github.com/sorohimm/uacs-store-back/pkg/product"

func NewProductFromRequest(req *product.CreateProductRequest) *Product {
	info := make([]*ProductInfo, 0, len(req.Info))
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

func (o *Product) ToAPIResponse() *product.ProductResponse {
	info := make([]*product.ProductInfo, 0, len(o.Info))
	for _, el := range o.Info {
		info = append(info, el.ToAPI())
	}
	return &product.ProductResponse{
		Id:    o.ID,
		Name:  o.Name,
		Price: o.Price,
		Img:   o.Img,
		Info:  info,
	}
}

type Products []*Product

func (o Products) ToAPIResponse() *product.AllProductsResponse {
	products := make([]*product.ProductResponse, 0, len(o))
	for _, el := range o {
		products = append(products, el.ToAPIResponse())
	}

	return &product.AllProductsResponse{Products: products}
}
