// Package dto TODO
package dto

import "github.com/sorohimm/uacs-store-back/pkg/product"

func NewProductInfosFromAPI(info []*product.ProductInfo) []*ProductInfo {
	to := make([]*ProductInfo, 0, len(info))
	for _, el := range info {
		to = append(to, NewProductInfoFromAPI(el))
	}
	return to
}

func NewProductInfoFromAPI(info *product.ProductInfo) *ProductInfo {
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

func (o ProductInfo) ToAPI() *product.ProductInfo {
	return &product.ProductInfo{
		ProductId:   o.ProductID,
		Title:       o.Title,
		Description: o.Description,
	}
}
