package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

type Config struct {
	DB     *sqlx.DB
	Logger *logrus.Logger
}

type mysqlCartRepository struct {
	db *sqlx.DB
}

func NewMysqlCartRepository(c Config) *mysqlCartRepository {
	log = c.Logger
	return &mysqlCartRepository{db: c.DB}
}
