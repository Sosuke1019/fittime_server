package router

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ponyo-E/fittime_server/model"
)

type ReqLog struct {
	MenuId uuid.UUID `json: "menuId"`
}

func AddLogHandler(c echo.Context) error {

	userId, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "uuidエラー")
	}

	var req ReqLog

	err = c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "バインドエラー")
	}

	err = model.AddLog(userId, req.MenuId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	return c.NoContent((http.StatusOK))
}
