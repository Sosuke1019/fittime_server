package router

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ponyo-E/fittime_server/model"
)

func AddLogHandler(c echo.Context) error {

	userId, err := uuid.Parse(c.Param("userId"))
	menuId, err := uuid.Parse(c.Param("menuId"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "uuidエラー")
	}

	err = model.AddLog(userId, menuId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	return c.NoContent(http.StatusOK)
}
