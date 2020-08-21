package service

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func (a *authService) Login(username string, password string) (success bool, err error) {
	hashed, err := a.authRepo.GetPassword(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	if !checkPasswordHashed(password, hashed) {
		log.WithFields(logrus.Fields{
			"username": username,
		}).Info("incorrect credential")
		return false, nil
	}
	return true, nil
}

func (a *authService) GetPosID(username string) (posID int, err error) {
	return a.authRepo.GetPosID(username)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHashed(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
