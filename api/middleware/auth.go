package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var key string

func InitAuth(signKey string) {
	key = signKey
}

func Auth() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(key),
		Skipper:       middleware.DefaultSkipper,
		SigningMethod: middleware.AlgorithmHS256,
		ContextKey:    "pos-id",
		TokenLookup:   "header:" + echo.HeaderAuthorization,
		AuthScheme:    "Bearer",
		Claims:        jwt.MapClaims{},
	})
}

func SetPosID2(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("pos-id", 1)
		return next(c)
	}
}
