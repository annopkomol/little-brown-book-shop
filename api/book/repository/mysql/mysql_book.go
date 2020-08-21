package mysql

import (
	"lbbs-service/domain"
)

func (r *mysqlBookRepository) GetAllBooks() (books []domain.Book, err error) {
	query := "SELECT id, title, cover, price FROM books;"
	err = r.db.Select(&books, query)
	if err != nil {
		log.Printf("couldn't query books the database: %s", err)
	}
	return
}
