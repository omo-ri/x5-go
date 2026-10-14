package order

import "errors"

type Item struct {
	Name         string
	PriceKopecks int64
}

type Order struct {
	ID     int
	Items  []Item
	Status string
}

const (
	StatusNew       = "new"
	StatusPaid      = "paid"
	StatusShipped   = "shipped"
	StatusDelivered = "delivered"
	StatusCancelled = "cancelled"
)

var StatusNames = map[string]string{
	StatusNew:       "Новый",
	StatusPaid:      "Оплачен",
	StatusShipped:   "Отправлен",
	StatusDelivered: "Доставлен",
	StatusCancelled: "Отменён",
}

var ErrInvalidOrderID = errors.New("orderID should be positive")
var ErrNilOrder = errors.New("order is nil")
var ErrEmptyItemName = errors.New("item name is empty")
var ErrNegativeItemPrice = errors.New("item price is negative")

func NewOrder(id int) (*Order, error) {
	if id <= 0 {
		return nil, ErrInvalidOrderID
	}
	return &Order{
		ID:     id,
		Status: StatusNew,
	}, nil
}

func AddItem(order *Order, item Item) error {
	if order == nil {
		return ErrNilOrder
	}

	if item.Name == "" {
		return ErrEmptyItemName
	}

	if item.PriceKopecks < 0 {
		return ErrNegativeItemPrice
	}
	order.Items = append(order.Items, item)
	return nil
}

func Total(order Order) int64 {
	var result int64
	for _, i := range order.Items {
		result += i.PriceKopecks
	}
	return result
}
