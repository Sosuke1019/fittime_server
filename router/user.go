package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ponyo-E/fittime_server/model"
	"net/http"
	"net/mail"
)

type ReqCreateUser struct {
	Username string `json:"username"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

func CreateUserHandler(c echo.Context) error {

	var req ReqCreateUser

	err := c.Bind(&req)

	if err != nil {
		//バインドが間違えるとエラーを出す
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	_, err = mail.ParseAddress(req.Mail)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	err = model.CreateUser(req.Username, req.Mail, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	return c.NoContent(http.StatusOK)
}
