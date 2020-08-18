package domain

import "github.com/shopspring/decimal"

type Discount struct {
	Message string          `json:"message"`
	Amount  decimal.Decimal `json:"amount"`
}

func (d *Discount) HasDiscount() bool {
	return d.Message != "" || d.Amount != decimal.Decimal{}
}

type Discounts struct {
	All []Discount
}

func (d *Discounts) TotalDiscount() decimal.Decimal {
	total := decimal.NewFromInt(0)
	for _, dis := range d.All {
		total = total.Add(dis.Amount)
	}
	return total
}

type DiscountService interface {
	CheckDiscount(cart Cart) Discounts
}
