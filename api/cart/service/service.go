package service

import (
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"lbbs-service/domain"
	"net/http"
)

func (s *service) GetCartIDFromPos(posID int) (cartID int, err error) {
	cart, err := s.cartRepo.FindCartByPosID(posID)
	if err != nil {
		return 0, err
	}
	return cart.ID, nil
}

func (s *service) GetCart(posID int) (cart domain.Cart, err error) {
	cart, err = s.cartRepo.FindCartByPosID(posID)
	if err == sql.ErrNoRows {
		return s.cartRepo.CreateNewCart(posID)
	}
	return cart, err
}

func (s *service) AddBookToCart(cartID, bookID int) (remaining int, err error) {
	order, err := s.cartRepo.FindOrder(cartID, bookID)
	if err == sql.ErrNoRows {
		order, err := s.cartRepo.CreateNewOrder(cartID, bookID)
		if err != nil {
			return 0, err
		}
		return order.Qty, nil
	}
	if err != nil {
		return 0, err
	}
	order.Qty++
	if err := s.cartRepo.UpdateOrderQty(order.ID, order.Qty); err != nil {
		return 0, err
	}
	return order.Qty, nil
}

func (s *service) RemoveBookFromCart(cartID, bookID, qty int) (remaining int, err error) {
	order, err := s.cartRepo.FindOrder(cartID, bookID)
	if err != nil {
		return 0, err
	}
	if order.Qty < qty {
		err := fmt.Errorf("invalid number of books")
		log.WithFields(logrus.Fields{
			"have":    order.Qty,
			"remove":  qty,
			"cart_id": cartID,
			"book_id": bookID,
		}).Error(err)
		return 0, domain.NewErrorWithConfig(http.StatusUnprocessableEntity, err.Error())
	}
	order.Qty -= qty
	if order.Qty == 0 {
		return 0, s.cartRepo.DeleteOrder(order.ID)
	}
	return order.Qty, s.cartRepo.UpdateOrderQty(order.ID, order.Qty)
}

func (s *service) Checkout(cartID int, cash decimal.Decimal) (change decimal.Decimal, err error) {
	cart, err := s.cartRepo.FindCartByID(cartID)
	if err != nil {
		return
	}
	totalAmount := cart.TotalAmount()
	discounts := s.discountService.CheckDiscount(cart)
	netAmount := totalAmount.Sub(discounts.TotalDiscount())
	if netAmount.LessThan(decimal.NewFromInt(0)) {
		netAmount = decimal.NewFromInt(0)
	}
	if cash.LessThan(netAmount) {
		err := fmt.Errorf("not enough cash")
		log.WithFields(logrus.Fields{
			"have":        cash,
			"total_price": totalAmount,
			"cart_id":     cartID,
		}).Error(err)
		return change, domain.NewErrorWithConfig(http.StatusUnprocessableEntity, err.Error())
	}
	change = cash.Sub(netAmount)
	return change, s.cartRepo.FlushCart(cart)
}
