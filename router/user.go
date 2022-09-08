package router

import (
	"fmt"
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

	fmt.Println(req)

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

type ResUser struct {
	ID      uuid.UUID
	Name    string
	Profile string
	Point   int
	Level   int
	Status  string
}

func GetUserHandler(c echo.Context) error {
	userId, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	user, err := model.GetUser(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	level, status := model.GetLevelAndStatus(user.Point)
	// user型　から ResUser型に変換
	resUser := ResUser{
		ID:      user.ID,
		Name:    user.Name,
		Profile: user.Profile,
		Point:   user.Point,
		Level:   level,
		Status:  status,
	}

	return c.JSON(http.StatusOK, resUser)
}

type ReqProfile struct {
	Profile string `json:"profile"`
}

func UpdateProfileHandler(c echo.Context) error {

	userId, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	var req ReqProfile

	err = c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	err = model.UpdateProfile(userId, req.Profile)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	return c.NoContent(http.StatusOK)
}

type ReqName struct {
	Name string `json:"name"`
}

func UpdateUsernameHandler(c echo.Context) error {

	userId, err := uuid.Parse((c.Param("userId")))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	var req ReqName

	err = c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	err = model.UpdateName(userId, req.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	return c.NoContent(http.StatusOK)
}
