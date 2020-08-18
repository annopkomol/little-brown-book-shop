package service

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"lbbs-service/domain"
)

var log *logrus.Logger

type authService struct {
	repo domain.AuthRepository
}

func NewAuthService(repo domain.AuthRepository, logger *logrus.Logger) *authService {
	log = logger
	return &authService{repo: repo}
}

func (a *authService) Login(username string, password string) (success bool, err error) {
	hashed, err := a.repo.GetPassword(username)
	if err != nil {
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
	return a.repo.GetPosID(username)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHashed(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
