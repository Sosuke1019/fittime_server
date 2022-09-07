package router

import (
	"github.com/google/uuid"

	"net/http"
	"net/mail"

	"github.com/labstack/echo/v4"
	"github.com/ponyo-E/fittime_server/model"
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

type ReqProfile struct {
	Profile string `json:"profile"`
}

func AddProfileHandler(c echo.Context) error {

	userId, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	var req ReqProfile

	err = c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	err = model.AddProfile(userId, req.Profile)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	return c.NoContent(http.StatusOK)
}

type ReqName struct {
	Name string `json:"name"`
}

func AddNameHandler(c echo.Context) error {

	userId, err := uuid.Parse((c.Param("userId")))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	var req ReqName

	err = c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	err = model.AddName(userId, req.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	return c.NoContent((http.StatusOK))
}
