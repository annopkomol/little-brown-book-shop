package main

import (
	"context"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	echomdw "github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	authrepository "lbbs-service/auth/repository/mysql"
	authservice "lbbs-service/auth/service"
	authhandler "lbbs-service/auth/transport/http"
	bookrepository "lbbs-service/book/repository/mysql"
	bookservice "lbbs-service/book/service"
	bookhandler "lbbs-service/book/transport/http"
	cartrepository "lbbs-service/cart/repository/mysql"
	cartservice "lbbs-service/cart/service"
	carthandler "lbbs-service/cart/transport/http"
	"lbbs-service/middleware"
	"lbbs-service/route"
	"lbbs-service/util"
	"lbbs-service/util/response"
	"os"
	"os/signal"
	"time"
)

func main() {

	//env
	var (
		//env := os.Getenv("ENVIRONMENT")
		signKey = os.Getenv("SIGNING_KEY")
		port    = os.Getenv("PORT")
		dbConn  = os.Getenv("DB_CONN")
	)
	// Echo
	e := echo.New()
	e.HideBanner = true

	// Log
	log := logrus.New()
	log.SetReportCaller(true)

	// Validator
	eng := en.New()
	uni := ut.New(eng, eng)
	trans, found := uni.GetTranslator("en")
	if !found {
		log.Warn("couldn't get english translator for validator")
	}
	validate := validator.New()
	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		log.Panic(err)
	}
	response.Init(trans)
	e.Validator = &echoValidator{validator: validate}

	// Middleware
	e.Pre(echomdw.RemoveTrailingSlash())

	e.Use(echomdw.Logger())
	e.Use(echomdw.Recover())
	e.Use(echomdw.CORS())
	//e.Use(echomdw.CSRF())

	// Auth
	middleware.InitAuth(signKey)
	util.InitAuth(signKey)

	// Database
	db, err := sqlx.Connect("mysql", dbConn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	// setting shorter life time can affect connection being killed frequently.
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(15 * time.Minute)

	// Dependency Injection
	bookRepo := bookrepository.NewMysqlBookRepository(db, log)
	bookSvc := bookservice.NewBookService(bookRepo, log)
	bookHdlr := bookhandler.NewBookHandler(bookSvc)

	cartRepo := cartrepository.NewMysqlCartRepository(db, log)
	cartSvc := cartservice.NewCartService(cartRepo, log)
	cartHdlr := carthandler.NewCartHandler(cartSvc)

	authRepo := authrepository.NewMysqlAuthRepository(db, log)
	authSvc := authservice.NewAuthService(authRepo, log)
	authHdlr := authhandler.NewAuthHandler(authSvc)

	route.Init(route.Config{
		Echo: e,
		Book: bookHdlr,
		Cart: cartHdlr,
		Auth: authHdlr,
	})

	// Start server
	go func() {
		util.PrintBanner()
		address, err := util.GetServerAddress(port)
		if err != nil {
			log.Panic(err)
		}
		if err := e.Start(address); err != nil {
			log.Info("shutting down the server")
		}
	}()
	sigWatcher := make(chan os.Signal)
	signal.Notify(sigWatcher, os.Interrupt, os.Kill)
	<-sigWatcher
	log.Println("catch signal, performing gracefully shutdown...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Panic(err)
	}
}

type echoValidator struct {
	validator *validator.Validate
}

func (v *echoValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
