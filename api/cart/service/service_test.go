package service

import (
	"database/sql"
	"github.com/shopspring/decimal"
	discountservice "lbbs-service/discount/service"
	"lbbs-service/domain"
	"lbbs-service/domain/mocks"
	"testing"
)

func Test_service_AddBookToCart(t *testing.T) {
	type args struct {
		cartID int
		bookID int
	}
	tests := []struct {
		name          string
		args          args
		stub          func(r *mocks.CartRepository, a args)
		wantRemaining int
		wantErr       bool
	}{
		{
			name: "1 + 1 = 2",
			args: args{
				cartID: 1,
				bookID: 1,
			},
			stub: func(m *mocks.CartRepository, a args) {
				order := domain.Order{
					ID:   1,
					Book: domain.Book{},
					Qty:  1,
				}
				m.On("FindOrder", a.cartID, a.bookID).Return(order, nil)
				m.On("UpdateOrderQty", order.ID, 2).Return(nil)
			},
			wantRemaining: 2,
			wantErr:       false,
		},
		{
			name: "0 + 1 = 1",
			args: args{
				cartID: 1,
				bookID: 1,
			},
			stub: func(m *mocks.CartRepository, a args) {
				order := domain.Order{
					ID:   1,
					Book: domain.Book{},
					Qty:  1,
				}
				m.On("FindOrder", a.cartID, a.bookID).Return(domain.Order{}, sql.ErrNoRows)
				m.On("CreateNewOrder", a.cartID, a.bookID).Return(order, nil)
				m.On("UpdateOrderQty", order.ID, 1).Return(nil)
			},
			wantRemaining: 1,
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(mocks.CartRepository)
			tt.stub(repo, tt.args)
			s := &service{
				cartRepo: repo,
			}
			gotRemaining, err := s.AddBookToCart(tt.args.cartID, tt.args.bookID)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddBookToCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRemaining != tt.wantRemaining {
				t.Errorf("AddBookToCart() gotRemaining = %v, want %v", gotRemaining, tt.wantRemaining)
			}
		})
	}
}

func Test_service_Checkout(t *testing.T) {

	type args struct {
		cartID int
		cash   decimal.Decimal
	}
	tests := []struct {
		name       string
		args       args
		stub       func(r *mocks.CartRepository, a args)
		wantChange decimal.Decimal
		wantErr    bool
	}{
		{
			name: "success",
			args: args{
				cartID: 1,
				cash:   decimal.NewFromInt(1300),
			},
			stub: func(r *mocks.CartRepository, a args) {
				cart := domain.Cart{
					ID:            0,
					PosTerminalID: 0,
					Orders: []domain.Order{{
						ID: 1,
						Book: domain.Book{
							ID:    1,
							Title: "Harry Potter and the Philosopher's Stone (I)",
							Price: decimal.NewFromInt(350),
						},
						Qty: 2,
					}, {
						ID: 2,
						Book: domain.Book{
							ID:    2,
							Title: "Harry Potter and the Chamber of Secrets (II)",
							Price: decimal.NewFromInt(350),
						},
						Qty: 1,
					}, {
						ID: 3,
						Book: domain.Book{
							ID:    2,
							Title: "some book",
							Price: decimal.NewFromInt(260),
						},
						Qty: 1,
					}},
				}
				r.On("FindCartByID", a.cartID).Return(cart, nil)
				r.On("FlushCart", cart).Return(nil)
			},
			wantChange: decimal.NewFromInt(60),
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(mocks.CartRepository)
			tt.stub(repo, tt.args)
			s := &service{
				cartRepo:        repo,
				discountService: discountservice.NewDiscountService(),
			}
			gotChange, err := s.Checkout(tt.args.cartID, tt.args.cash)
			if (err != nil) != tt.wantErr {
				t.Errorf("Checkout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !gotChange.Equal(tt.wantChange) {
				t.Errorf("Checkout() gotChange = %v, want %v", gotChange, tt.wantChange)
			}
		})
	}
}

func Test_service_RemoveBookFromCart(t *testing.T) {
	type args struct {
		cartID int
		bookID int
		qty    int
	}
	tests := []struct {
		name          string
		args          args
		stub          func(r *mocks.CartRepository, a args)
		wantRemaining int
		wantErr       bool
	}{
		{
			name: "2 - 1 = 1",
			args: args{
				cartID: 1,
				bookID: 1,
				qty:    1,
			},
			stub: func(m *mocks.CartRepository, a args) {
				order := domain.Order{
					ID:   1,
					Book: domain.Book{},
					Qty:  2,
				}
				m.On("FindOrder", a.cartID, a.bookID).Return(order, nil)
				m.On("UpdateOrderQty", order.ID, 1).Return(nil)
			},
			wantRemaining: 1,
			wantErr:       false,
		},
		{
			name: "1 - 1 = 0",
			args: args{
				cartID: 1,
				bookID: 1,
				qty:    1,
			},
			stub: func(m *mocks.CartRepository, a args) {
				order := domain.Order{
					ID:   1,
					Book: domain.Book{},
					Qty:  1,
				}
				m.On("FindOrder", a.cartID, a.bookID).Return(order, nil)
				m.On("DeleteOrder", order.ID).Return(nil)
			},
			wantRemaining: 0,
			wantErr:       false,
		},
		{
			name: "0 - 1 = error",
			args: args{
				cartID: 1,
				bookID: 1,
				qty:    1,
			},
			stub: func(m *mocks.CartRepository, a args) {
				m.On("FindOrder", a.cartID, a.bookID).Return(domain.Order{}, sql.ErrNoRows)
			},
			wantRemaining: 0,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(mocks.CartRepository)
			tt.stub(repo, tt.args)
			s := &service{
				cartRepo: repo,
			}
			gotRemaining, err := s.RemoveBookFromCart(tt.args.cartID, tt.args.bookID, tt.args.qty)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveBookFromCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRemaining != tt.wantRemaining {
				t.Errorf("RemoveBookFromCart() gotRemaining = %v, want %v", gotRemaining, tt.wantRemaining)
			}
		})
	}
}
