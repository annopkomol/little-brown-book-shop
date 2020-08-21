package service

import (
	"github.com/shopspring/decimal"
	"lbbs-service/domain"
)

type DiscountService struct {
}

func NewDiscountService() *DiscountService {
	return &DiscountService{}
}

func (d *DiscountService) CheckDiscount(cart domain.Cart) (discounts domain.Discounts) {
	harryDiscount := uniqueHarryPotterDiscount(cart)
	if !harryDiscount.Amount.IsZero() {
		discounts.All = append(discounts.All, harryDiscount)
	}
	return
}

func discountAmount(totalAmount decimal.Decimal, discountPercent int) decimal.Decimal {
	discount := totalAmount.Mul(decimal.NewFromInt(int64(discountPercent))).Div(decimal.NewFromInt(100))
	if discount.GreaterThan(totalAmount) {
		return totalAmount
	}
	return discount
}

func uniqueHarryPotterDiscount(cart domain.Cart) domain.Discount {
	var (
		unique, currentDiscount int
		currentRule             uniqueHarryPotterRule
		totalPrice              = decimal.NewFromInt(0)
	)
	isHarryPotter := func(title string) bool {
		for _, series := range harryPotterSeries {
			if title == series {
				return true
			}
		}
		return false
	}

	for _, order := range cart.Orders {
		isHP := isHarryPotter(order.Book.Title)
		if isHP {
			totalPrice = decimal.Sum(totalPrice, order.Book.Price)
			unique++
		}
	}
	for _, rule := range uniqueHarryPotterRules {
		if unique >= rule.Qty && rule.PercentDiscount > currentDiscount {
			currentRule = rule
		}
	}

	if currentRule.PercentDiscount == 0 {
		return domain.Discount{}
	}
	return domain.Discount{
		Message: currentRule.Msg,
		Amount:  discountAmount(totalPrice, currentRule.PercentDiscount),
	}
}
