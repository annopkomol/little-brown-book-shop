package domain

import "github.com/shopspring/decimal"

type Book struct {
	ID    int             `json:"id" db:"id"`
	Title string          `json:"title" db:"title"`
	Cover string          `json:"cover" db:"cover"`
	Price decimal.Decimal `json:"price" db:"price"`
}

type BookService interface {
	GetAllBooks() ([]Book, error)
}

type BookRepository interface {
	GetAllBooks() ([]Book, error)
}
