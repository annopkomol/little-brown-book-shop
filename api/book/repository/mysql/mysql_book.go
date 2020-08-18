package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"lbbs-service/domain"
)

var log = logrus.New()

type mysqlBookRepository struct {
	db *sqlx.DB
}

func NewMysqlBookRepository(db *sqlx.DB, logger *logrus.Logger) *mysqlBookRepository {
	log = logger
	return &mysqlBookRepository{db: db}
}

func (r *mysqlBookRepository) GetAllBooks() (books []domain.Book, err error) {
	query := "SELECT id, tittle, cover, price FROM books;"
	err = r.db.Select(&books, query)
	if err != nil {
		log.Printf("couldn't query books the database: %s", err)
	}
	return
}
