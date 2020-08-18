package route

import (
	"github.com/labstack/echo"
	authhttp "lbbs-service/auth/transport/http"
	bookhttp "lbbs-service/book/transport/http"
	carthttp "lbbs-service/cart/transport/http"
)

type Config struct {
	Echo *echo.Echo
	Book *bookhttp.BookHandler
	Cart *carthttp.CartHandler
	Auth *authhttp.AuthHandler
}

func initConfig(c Config) {
	e = c.Echo
	book = c.Book
	cart = c.Cart
	auth = c.Auth
}
