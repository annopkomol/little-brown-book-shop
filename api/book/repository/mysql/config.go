package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type Config struct {
	DB     *sqlx.DB
	Logger *logrus.Logger
}

type mysqlBookRepository struct {
	db *sqlx.DB
}

func NewMysqlBookRepository(c Config) *mysqlBookRepository {
	log = c.Logger
	return &mysqlBookRepository{db: c.DB}
}
