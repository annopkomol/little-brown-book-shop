package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

type mysqlAuthRepository struct {
	db *sqlx.DB
}

func NewMysqlAuthRepository(db *sqlx.DB, logger *logrus.Logger) *mysqlAuthRepository {
	log = logger
	return &mysqlAuthRepository{db: db}
}

func (a *mysqlAuthRepository) GetPassword(username string) (hashed string, err error) {
	err = a.db.QueryRow(
		"SELECT password FROM pos_terminals WHERE username = ?;",
		username).Scan(&hashed)
	if err != nil {
		log.Error(err)
		return "", err
	}
	return hashed, nil
}

func (a *mysqlAuthRepository) GetPosID(username string) (posID int, err error) {
	err = a.db.QueryRow(
		"SELECT id FROM pos_terminals WHERE username = ?;",
		username).Scan(&posID)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return posID, nil
}
