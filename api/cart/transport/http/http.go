package http

import (
	"github.com/labstack/echo"
	"github.com/shopspring/decimal"
	"lbbs-service/domain"
	"lbbs-service/util"
	res "lbbs-service/util/response"
	"strconv"
)

type CartHandler struct {
	service  domain.CartService
	discount domain.DiscountService
}

func NewCartHandler(service domain.CartService, discount domain.DiscountService) *CartHandler {
	return &CartHandler{service: service, discount: discount}
}

func (h *CartHandler) GetCart(c echo.Context) error {
	type discountFixedStr struct {
		Message string `json:"message"`
		Amount  string `json:"amount"`
	}
	resp := struct {
		Cart          domain.Cart        `json:"cart"`
		Discount      []discountFixedStr `json:"discounts"`
		TotalDiscount string             `json:"total_discount"`
		Net           string             `json:"net_amount"`
	}{}

	posID := util.GetToken(c)
	cart, err := h.service.GetCart(posID)
	if err != nil {
		return res.Error(c, err)
	}
	discounts := h.discount.CheckDiscount(cart)
	resp.Cart = cart
	for _, d := range discounts.All {
		resp.Discount = append(resp.Discount, discountFixedStr{
			Message: d.Message,
			Amount:  d.Amount.StringFixedBank(2),
		})
	}
	resp.TotalDiscount = discounts.TotalDiscount().StringFixedBank(2)
	resp.Net = cart.TotalAmount().Sub(discounts.TotalDiscount()).StringFixedBank(2)
	return res.Success(c, resp)
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
	posID := util.GetToken(c)
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
	posID := util.GetToken(c)
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
	posID := util.GetToken(c)
	cartID, err := h.service.GetCartIDFromPos(posID)
	if err != nil {
		return res.Error(c, err)
	}
	change, err := h.service.Checkout(cartID, cash)
	if err != nil {
		return res.Error(c, err)
	}
	return res.Success(c, map[string]string{"change": change.StringFixedBank(2)})
}
