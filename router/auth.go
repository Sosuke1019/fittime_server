package router

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ponyo-E/fittime_server/model"
	"net/http"
)

type LoginRequest struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type ResLogin struct {
	UserID uuid.UUID `json:"userID"`
}

func LoginHandler(c echo.Context) error {
	var req LoginRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	user, err := model.CheckPassword(req.Mail, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}
	res := ResLogin{
		UserID: user.ID,
	}
	return c.JSON(http.StatusOK, res)
}
