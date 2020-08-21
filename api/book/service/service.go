package service

import (
	"lbbs-service/domain"
)

func (b *BookService) GetAllBooks() ([]domain.Book, error) {
	return b.bookRepo.GetAllBooks()
}
