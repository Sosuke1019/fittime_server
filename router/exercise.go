package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ponyo-E/fittime_server/model"
	"net/http"
)

func GetExerciseHandler(c echo.Context) error {
	exercises, err := model.GetExercises()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	return c.JSON(http.StatusOK, exercises)
}
