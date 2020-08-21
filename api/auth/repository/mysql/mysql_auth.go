package mysql

import (
	"database/sql"
)

func (a *mysqlAuthRepository) GetPassword(username string) (hashed string, err error) {
	err = a.db.QueryRow(
		"SELECT password FROM pos_terminals WHERE username = ?;",
		username).Scan(&hashed)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Error(err)
		}
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
