package http

import (
	"github.com/labstack/echo"
	"lbbs-service/domain"
	"net/http"
)

type BookHandler struct {
	service domain.BookService
}

func NewBookHandler(service domain.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (b *BookHandler) GetBookList(c echo.Context) error {
	books, err := b.service.GetAllBooks()
	if err != nil {
		return c.String(500, err.Error())
	}
	return c.JSON(http.StatusOK, books)
}
