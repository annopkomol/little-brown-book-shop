package domain

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestCart_TotalAmount(t *testing.T) {
	type fields struct {
		ID            int
		PosTerminalID int
		Orders        []Order
	}
	tests := []struct {
		name   string
		fields fields
		want   decimal.Decimal
	}{
		{
			name: "100 + 100.2 + 100.3 = 300.5",
			fields: fields{
				Orders: []Order{{
					Book: Book{
						Price: decimal.NewFromFloat(100),
					},
					Qty: 1,
				}, {
					Book: Book{
						Price: decimal.NewFromFloat(100.2),
					},
					Qty: 1,
				}, {
					Book: Book{
						Price: decimal.NewFromFloat(100.3),
					},
					Qty: 1,
				}},
			},
			want: decimal.NewFromFloat(300.5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cart{
				ID:            tt.fields.ID,
				PosTerminalID: tt.fields.PosTerminalID,
				Orders:        tt.fields.Orders,
			}
			if got := c.TotalAmount(); !got.Equal(tt.want) {
				t.Errorf("TotalAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_TotalPrice(t *testing.T) {
	type fields struct {
		ID   int
		Book Book
		Qty  int
	}
	tests := []struct {
		name   string
		fields fields
		want   decimal.Decimal
	}{
		{
			name: "100 * 3 = 300",
			fields: fields{
				Book: Book{
					Price: decimal.NewFromInt(100),
				},
				Qty: 3,
			},
			want: decimal.NewFromInt(300),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Order{
				ID:   tt.fields.ID,
				Book: tt.fields.Book,
				Qty:  tt.fields.Qty,
			}
			if got := o.TotalPrice(); !got.Equal(tt.want) {
				t.Errorf("TotalPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
