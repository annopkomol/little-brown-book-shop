package http

import (
	"github.com/labstack/echo"
	"lbbs-service/domain"
	"lbbs-service/util"
	res "lbbs-service/util/response"
)

type AuthHandler struct {
	service domain.AuthService
}

func NewAuthHandler(service domain.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (a *AuthHandler) Login(c echo.Context) error {
	req := struct {
		Username string `validate:"required"`
		Password string `validate:"required"`
	}{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}
	if err := c.Validate(req); err != nil {
		return res.Error(c, err)
	}
	success, err := a.service.Login(req.Username, req.Password)
	if err != nil {
		return res.Error(c, err)
	}
	if !success {
		return res.Error(c, domain.NewErrorWithConfig(403, "credential is incorrect"))
	}
	posID, err := a.service.GetPosID(req.Username)
	if err != nil {
		return res.Error(c, err)
	}
	token, err := util.CreateToken(posID)
	if err != nil {
		return res.Error(c, err)
	}
	return res.Success(c, map[string]interface{}{
		"access_token": token,
	})
}
