package cart

import (
	"github.com/sorohimm/uacs-store-back/pkg/api"
)

func NewItemFromApi(item *api.CartItem) *Item {
	return &Item{
		CartID:    item.CartId,
		ProductID: item.ProductId,
		Quantity:  item.Quantity,
	}
}

type Item struct {
	ID        int64
	CartID    int64
	ProductID int64
	Sku       int64
	Name      string
	Price     float32
	Quantity  int64
}

func (o *Item) SetID(id int64) *Item {
	o.ID = id
	return o
}

func NewCart() *Cart {
	return &Cart{}
}

type Cart struct {
	ID    int64
	Items []*Item
}

type Customer struct {
	Id        int64
	Email     string
	FirstName string
	LastName  string
	Phone     string
}

type CartInfo struct {
	Id       int64
	Customer *Customer
	Items    []*Item
}

func (o *CartInfo) ToAPI() *api.CartInfo {
	return nil // TODO: implement
}

func (o *Cart) SetID(id int64) *Cart {
	o.ID = id
	return o
}

func (o *Cart) SetItems(items []*Item) *Cart {
	o.Items = items
	return o
}
