package service

import (
	"github.com/sirupsen/logrus"
	"lbbs-service/domain"
)

var log *logrus.Logger

type Config struct {
	AuthRepository domain.AuthRepository
	Logger         *logrus.Logger
}

type authService struct {
	authRepo domain.AuthRepository
}

func NewAuthService(c Config) *authService {
	log = c.Logger
	return &authService{authRepo: c.AuthRepository}
}
