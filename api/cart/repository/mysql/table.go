package mysql

import "github.com/shopspring/decimal"

type cartTable struct {
	ID            int `db:"id"`
	PosTerminalID int `db:"pos_terminal_id"`
}

type orderTable struct {
	ID     int `db:"id"`
	CartID int `db:"cart_id"`
	BookID int `db:"book_id"`
	Qty    int `db:"qty"`
}

type orderWithBookTable struct {
	orderTable
	Title string          `db:"title"`
	Cover string          `db:"cover"`
	Price decimal.Decimal `db:"price"`
}
