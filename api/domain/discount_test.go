package domain

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestDiscount_HasDiscount(t *testing.T) {
	type fields struct {
		Message string
		Amount  decimal.Decimal
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "no discount",
			fields: fields{
				Message: "",
				Amount:  decimal.Decimal{},
			},
			want: false,
		}, {
			name: "has discount",
			fields: fields{
				Message: "123123",
				Amount:  decimal.NewFromInt(1),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Discount{
				Message: tt.fields.Message,
				Amount:  tt.fields.Amount,
			}
			if got := d.HasDiscount(); got != tt.want {
				t.Errorf("HasDiscount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscounts_TotalDiscount(t *testing.T) {
	type fields struct {
		All []Discount
	}
	tests := []struct {
		name   string
		fields fields
		want   decimal.Decimal
	}{
		{
			name: "100 + 150 = 250",
			fields: fields{
				All: []Discount{{
					Message: "asd",
					Amount:  decimal.NewFromInt(100),
				}, {
					Message: "1123",
					Amount:  decimal.NewFromInt(150),
				}},
			},
			want: decimal.NewFromInt(250),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Discounts{
				All: tt.fields.All,
			}
			if got := d.TotalDiscount(); !got.Equal(tt.want) {
				t.Errorf("TotalDiscount() = %v, want %v", got, tt.want)
			}
		})
	}
}
