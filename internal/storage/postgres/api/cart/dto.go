package cart

import "github.com/sorohimm/uacs-store-back/pkg/api"

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

func (o *Cart) SetID(id int64) *Cart {
	o.ID = id
	return o
}

func (o *Cart) SetItems(items []*Item) *Cart {
	o.Items = items
	return o
}
