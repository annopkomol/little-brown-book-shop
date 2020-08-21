package service

import (
	"github.com/sirupsen/logrus"
	"lbbs-service/domain"
)

var log = logrus.New()

type Config struct {
	CartRepository  domain.CartRepository
	DiscountService domain.DiscountService
	Logger          *logrus.Logger
}

type service struct {
	cartRepo        domain.CartRepository
	discountService domain.DiscountService
}

func NewCartService(c Config) *service {
	log = c.Logger
	return &service{cartRepo: c.CartRepository, discountService: c.DiscountService}
}
