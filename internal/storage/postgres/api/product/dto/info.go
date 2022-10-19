// Package dto TODO
package dto

import "github.com/sorohimm/uacs-store-back/pkg/api"

func NewProductInfosFromAPI(info []*api.ProductInfo) []*ProductInfo {
	to := make([]*ProductInfo, 0, len(info))
	for _, el := range info {
		to = append(to, NewProductInfoFromAPI(el))
	}
	return to
}

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
