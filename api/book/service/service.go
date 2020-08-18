package service

import (
	"github.com/sirupsen/logrus"
	"lbbs-service/domain"
)

var log *logrus.Logger

type BookService struct {
	repo domain.BookRepository
}

func NewBookService(book domain.BookRepository, logger *logrus.Logger) *BookService {
	log = logger
	return &BookService{repo: book}
}

func (b *BookService) GetAllBooks() ([]domain.Book, error) {
	return b.repo.GetAllBooks()
}
