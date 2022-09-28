package product

import "github.com/sorohimm/shop/pkg/api"

func NewProductInfoFromAPI(info *api.ProductInfo) *ProductInfo {
	return &ProductInfo{
		ProductID:   info.ProductId,
		Title:       info.Title,
		Description: info.Description,
	}
}

type ProductInfo struct {
	ProductID   int64
	Title       string
	Description string
}

func (o *ProductInfo) SetProductID(productID int64) *ProductInfo {
	o.ProductID = productID
	return o
}

func (o ProductInfo) ToAPI() *api.ProductInfo {
	return &api.ProductInfo{
		ProductId:   o.ProductID,
		Title:       o.Title,
		Description: o.Description,
	}
}

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
	return &api.ProductResponse{
		Id:    o.ID,
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
