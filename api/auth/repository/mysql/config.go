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

type mysqlAuthRepository struct {
	db *sqlx.DB
}

func NewMysqlAuthRepository(c Config) *mysqlAuthRepository {
	log = c.Logger
	return &mysqlAuthRepository{db: c.DB}
}
