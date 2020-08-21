package route

import (
	"github.com/labstack/echo"
	authhttp "lbbs-service/auth/transport/http"
	bookhttp "lbbs-service/book/transport/http"
	carthttp "lbbs-service/cart/transport/http"
	"lbbs-service/middleware"
)

var (
	e    *echo.Echo
	book *bookhttp.BookHandler
	cart *carthttp.CartHandler
	auth *authhttp.AuthHandler
)

func Init(c Config) {
	initConfig(c)

	// book
	bookRoute := e.Group("/books", middleware.Auth())
	bookRoute.GET("", book.GetBookList)

	// cart
	cartRoute := e.Group("/cart", middleware.Auth(), middleware.SetToken)
	cartRoute.GET("", cart.GetCart)
	cartRoute.POST("/:book-id", cart.AddBookToCart)
	cartRoute.DELETE("/:book-id", cart.RemoveBookToCart)
	cartRoute.POST("/checkout", cart.Checkout)

	// auth
	authRoute := e.Group("auth")
	authRoute.POST("/login", auth.Login)
}
