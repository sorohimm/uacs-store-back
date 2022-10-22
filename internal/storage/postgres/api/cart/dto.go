package cart

type Item struct {
	CartID    int64
	ProductID int64
	Quantity  int64
}

type Cart struct {
	ID    int64
	Items []*Item
}
