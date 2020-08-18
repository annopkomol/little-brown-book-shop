package response

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"lbbs-service/domain"
	"net/http"
)

var (
	trans ut.Translator
)

func Init(translator ut.Translator) {
	trans = translator
}

func Error(c echo.Context, err error) error {
	switch err := err.(type) {
	case validator.ValidationErrors:
		if len(err) != 0 {
			//TODO:: handle multiple validation errors
			e := err[0]
			return c.String(http.StatusBadRequest, e.Translate(trans))
		}
		return c.String(http.StatusBadRequest, "")
	case domain.Error:
		return c.String(err.GetHttpStatus(), err.Error())
	default:
		return c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}

func Success(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}
