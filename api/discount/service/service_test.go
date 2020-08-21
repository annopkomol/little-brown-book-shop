package service

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"lbbs-service/domain"
	"testing"
)

func TestDiscountService_CheckDiscount(t *testing.T) {
	type args struct {
		cart domain.Cart
	}
	tests := []struct {
		name          string
		args          args
		wantDiscounts domain.Discounts
	}{
		{
			name: "2 Harry 1 Fake :D",
			args: args{cart: domain.Cart{Orders: []domain.Order{{
				Book: domain.Book{
					ID:    0,
					Title: "Harry Potter and the Philosopher's Stone (I)",
					Cover: "",
					Price: decimal.NewFromInt(350),
				},
				Qty: 2,
			}, {
				Book: domain.Book{
					Title: "Harry Potter and the Chamber of Secrets (II)",
					Price: decimal.NewFromInt(350),
				},
				Qty: 1,
			}, {
				Book: domain.Book{
					Title: "I am a fake Harry Potter !! (XIII)",
					Price: decimal.NewFromInt(99999),
				},
				Qty: 1,
			},
			}}},
			wantDiscounts: domain.Discounts{
				All: []domain.Discount{{
					Message: "",
					Amount:  decimal.NewFromInt(70),
				}},
			},
		},
		{
			name: "no discount",
			args: args{
				cart: domain.Cart{
					Orders: []domain.Order{
						{
							Book: domain.Book{
								Title: "harry",
								Price: decimal.NewFromInt(123123),
							},
							Qty: 7,
						}, {
							Book: domain.Book{
								Title: "harry",
								Price: decimal.NewFromInt(123123),
							},
							Qty: 7,
						}, {
							Book: domain.Book{
								Title: "harry",
								Price: decimal.NewFromInt(123123),
							},
							Qty: 7,
						}, {
							Book: domain.Book{
								Title: "harry",
								Price: decimal.NewFromInt(123123),
							},
							Qty: 7,
						}, {
							Book: domain.Book{
								Title: "harry",
								Price: decimal.NewFromInt(123123),
							},
							Qty: 7,
						},
					},
				},
			},
			wantDiscounts: domain.Discounts{
				All: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DiscountService{}
			gotDiscounts := d.CheckDiscount(tt.args.cart)
			require.Equal(t, len(tt.wantDiscounts.All), len(gotDiscounts.All))
			for index, want := range tt.wantDiscounts.All {
				assert.True(t, want.Amount.Equal(gotDiscounts.All[index].Amount))
			}
		})
	}
}

func Test_discountAmount(t *testing.T) {
	type args struct {
		totalAmount     decimal.Decimal
		discountPercent int
	}
	tests := []struct {
		name string
		args args
		want decimal.Decimal
	}{
		{
			name: "10% of 1,000",
			args: args{
				totalAmount:     decimal.NewFromInt(1000),
				discountPercent: 10,
			},
			want: decimal.NewFromInt(100),
		}, {
			name: "10% of 9999.99",
			args: args{
				totalAmount:     decimal.NewFromFloat(9999.99),
				discountPercent: 10,
			},
			want: decimal.NewFromFloat(999.999),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := discountAmount(tt.args.totalAmount, tt.args.discountPercent); !got.Equal(tt.want) {
				t.Errorf("discountAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uniqueHarryPotterDiscount(t *testing.T) {
	type args struct {
		cart domain.Cart
	}
	tests := []struct {
		name string
		args args
		want domain.Discount
	}{
		{
			name: "2 Harry 1 Fake :D",
			args: args{
				cart: domain.Cart{
					Orders: []domain.Order{{
						Book: domain.Book{
							ID:    0,
							Title: "Harry Potter and the Philosopher's Stone (I)",
							Cover: "",
							Price: decimal.NewFromInt(350),
						},
						Qty: 3,
					}, {
						Book: domain.Book{
							Title: "Harry Potter and the Chamber of Secrets (II)",
							Price: decimal.NewFromInt(350),
						},
						Qty: 2,
					}, {
						Book: domain.Book{
							Title: "I am a fake Harry Potter !! (XIII)",
							Price: decimal.NewFromInt(99999),
						},
						Qty: 1,
					},
					},
				},
			},
			want: domain.Discount{
				Amount: decimal.NewFromInt(70),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueHarryPotterDiscount(tt.args.cart); !tt.want.Amount.Equal(got.Amount) {
				t.Errorf("uniqueHarryPotterDiscount() = %v, want %v", got.Amount, tt.want.Amount)
			}
		})
	}
}
