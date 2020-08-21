package domain

import (
	"github.com/shopspring/decimal"
)

type Cart struct {
	ID            int     `json:"id"`
	PosTerminalID int     `json:"pos_terminal_id"`
	Orders        []Order `json:"orders"`
}

func (c *Cart) TotalAmount() decimal.Decimal {
	var amount decimal.Decimal
	for _, o := range c.Orders {
		orderAmount := o.TotalPrice()
		amount = decimal.Sum(amount, orderAmount)
	}
	return amount
}

type Order struct {
	ID   int  `json:"id"`
	Book Book `json:"book"`
	Qty  int  `json:"qty"`
}

func (o *Order) TotalPrice() decimal.Decimal {
	return o.Book.Price.Mul(decimal.NewFromInt(int64(o.Qty)))
}

type CartService interface {
	GetCartIDFromPos(posID int) (cartID int, err error)
	GetCart(posID int) (cart Cart, err error)
	AddBookToCart(cartID, bookID int) (remaining int, err error)
	RemoveBookFromCart(cartID, bookID, qty int) (remaining int, err error)
	Checkout(cartID int, cash decimal.Decimal) (change decimal.Decimal, err error)
}

type CartRepository interface {
	FindCartByPosID(posID int) (Cart, error)
	FindCartByID(cartID int) (Cart, error)
	CreateNewCart(posID int) (Cart, error)
	FlushCart(cart Cart) error
	CountOrdersInCart(cartID int) (int, error)

	FindOrder(cartID, BookID int) (Order, error)
	CreateNewOrder(cartID, bookID int) (Order, error)
	UpdateOrderQty(orderID, qty int) error
	DeleteOrder(orderID int) error
}
