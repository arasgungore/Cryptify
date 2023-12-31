package order

import (
	"fmt"
	"sync"
)

// OrderType represents the type of order (Buy or Sell)
type OrderType int

const (
	Buy OrderType = iota
	Sell
)

// Order represents a trading order
type Order struct {
	ID       int
	UserID   int
	Base     string
	Quote    string
	Type     OrderType
	Quantity float64
	Price    float64
}

// OrderService handles order-related operations
type OrderService struct {
	orders      map[int]Order
	ordersMutex sync.Mutex
}

// NewOrderService creates a new instance of OrderService
func NewOrderService() *OrderService {
	return &OrderService{
		orders: make(map[int]Order),
	}
}

// PlaceOrder places a new order
func (o *OrderService) PlaceOrder(userID int, base, quote string, orderType OrderType, quantity, price float64) int {
	o.ordersMutex.Lock()
	defer o.ordersMutex.Unlock()

	orderID := len(o.orders) + 1
	newOrder := Order{
		ID:       orderID,
		UserID:   userID,
		Base:     base,
		Quote:    quote,
		Type:     orderType,
		Quantity: quantity,
		Price:    price,
	}

	o.orders[orderID] = newOrder
	fmt.Printf("Order placed: %v\n", newOrder)

	return orderID
}
