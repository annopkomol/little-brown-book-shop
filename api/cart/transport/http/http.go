package http

import (
	"github.com/labstack/echo"
	"github.com/shopspring/decimal"
	"lbbs-service/domain"
	res "lbbs-service/util/response"
	"strconv"
)

type CartHandler struct {
	service domain.CartService
}

func NewCartHandler(service domain.CartService) *CartHandler {
	return &CartHandler{service: service}
}

func (h *CartHandler) GetCart(c echo.Context) error {
	posID, _ := c.Get("pos-id").(int)
	cart, err := h.service.GetCart(posID)
	if err != nil {
		return res.Error(c, err)
	}
	return res.Success(c, cart)
}

func (h *CartHandler) AddBookToCart(c echo.Context) error {
	req := struct {
		BookID string `validate:"required,number"`
	}{
		BookID: c.Param("book-id"),
	}
	if err := c.Validate(req); err != nil {
		return res.Error(c, err)
	}
	posID := c.Get("pos-id").(int)
	cartID, err := h.service.GetCartIDFromPos(posID)
	if err != nil {
		return res.Error(c, err)
	}
	bookID, _ := strconv.Atoi(req.BookID)
	remaining, err := h.service.AddBookToCart(cartID, bookID)
	if err != nil {
		return res.Error(c, err)
	}
	return res.Success(c, map[string]int{"remaining": remaining})
}

func (h *CartHandler) RemoveBookToCart(c echo.Context) error {
	req := struct {
		BookID string `validate:"required,number"`
		Qty    string `validate:"omitempty,number"`
	}{
		BookID: c.Param("book-id"),
		Qty:    c.QueryParam("qty"),
	}
	if err := c.Validate(req); err != nil {
		return res.Error(c, err)
	}
	if req.Qty == "" {
		req.Qty = "1"
	}
	bookID, _ := strconv.Atoi(req.BookID)
	qty, err := strconv.Atoi(req.Qty)
	if err != nil {
		return res.Error(c, err)
	}
	posID := c.Get("pos-id").(int)
	cartID, err := h.service.GetCartIDFromPos(posID)
	if err != nil {
		return res.Error(c, err)
	}
	remaining, err := h.service.RemoveBookFromCart(cartID, bookID, qty)
	if err != nil {
		return res.Error(c, err)
	}
	return res.Success(c, map[string]int{"remaining": remaining})
}

func (h *CartHandler) Checkout(c echo.Context) error {
	req := struct {
		Cash string `validate:"required,numeric"`
	}{
		Cash: c.FormValue("cash"),
	}
	if err := c.Validate(req); err != nil {
		return res.Error(c, err)
	}
	cash, err := decimal.NewFromString(req.Cash)
	if err != nil {
		return res.Error(c, err)
	}
	posID := c.Get("pos-id").(int)
	cartID, err := h.service.GetCartIDFromPos(posID)
	if err != nil {
		return res.Error(c, err)
	}
	charge, err := h.service.Checkout(cartID, cash)
	if err != nil {
		return res.Error(c, err)
	}
	return res.Success(c, map[string]decimal.Decimal{"charge": charge})
}
