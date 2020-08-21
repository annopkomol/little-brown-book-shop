package service

import (
	"github.com/sirupsen/logrus"
	"lbbs-service/domain"
)

var log *logrus.Logger

type Config struct {
	BookRepository domain.BookRepository
	Logger         *logrus.Logger
}

type BookService struct {
	bookRepo domain.BookRepository
}

func NewBookService(c Config) *BookService {
	log = c.Logger
	return &BookService{bookRepo: c.BookRepository}
}
