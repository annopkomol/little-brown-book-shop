package service

import (
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"lbbs-service/domain"
	"net/http"
)

var log = logrus.New()

type service struct {
	repo domain.CartRepository
}

func (s *service) GetCartIDFromPos(posID int) (cartID int, err error) {
	cart, err := s.repo.FindCartByPosID(posID)
	if err != nil {
		return 0, err
	}
	return cart.ID, nil
}

func NewCartService(repo domain.CartRepository, logger *logrus.Logger) *service {
	log = logger
	return &service{repo: repo}
}

func (s *service) GetCart(posID int) (cart domain.Cart, err error) {
	cart, err = s.repo.FindCartByPosID(posID)
	if err == sql.ErrNoRows {
		return s.repo.CreateNewCart(posID)
	}
	return cart, err
}

func (s *service) AddBookToCart(cartID, bookID int) (remaining int, err error) {
	order, err := s.repo.FindOrder(cartID, bookID)
	if err == sql.ErrNoRows {
		order, err := s.repo.CreateNewOrder(cartID, bookID)
		if err != nil {
			return 0, err
		}
		return order.Qty, nil
	}
	if err != nil {
		return 0, err
	}
	order.Qty++
	if err := s.repo.UpdateOrderQty(order.ID, order.Qty); err != nil {
		return 0, err
	}
	return order.Qty, nil
}

func (s *service) RemoveBookFromCart(cartID, bookID, qty int) (remaining int, err error) {
	order, err := s.repo.FindOrder(cartID, bookID)
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
		return 0, s.repo.DeleteOrder(order.ID)
	}
	return order.Qty, s.repo.UpdateOrderQty(order.ID, order.Qty)
}

func (s *service) Checkout(cartID int, cash decimal.Decimal) (charge decimal.Decimal, err error) {
	cart, err := s.repo.FindCartByID(cartID)
	if err != nil {
		return
	}
	totalAmount := cart.TotalAmount()
	if cash.LessThan(totalAmount) {
		err := fmt.Errorf("not enough cash")
		log.WithFields(logrus.Fields{
			"have":        cash,
			"total_price": totalAmount,
			"cart_id":     cartID,
		}).Error(err)
		return charge, domain.NewErrorWithConfig(http.StatusUnprocessableEntity, err.Error())
	}
	charge = cash.Sub(totalAmount)
	return charge, s.repo.FlushCart(cart)
}
