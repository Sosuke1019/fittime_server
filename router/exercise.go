package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ponyo-E/fittime_server/model"
	"net/http"
)

type ResExercises struct {
	Exercises []model.Exercise `json:"exercises"`
}

func GetExerciseHandler(c echo.Context) error {
	exercises, err := model.GetExercises()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	resExercises := ResExercises{
		Exercises: exercises,
	}
	return c.JSON(http.StatusOK, resExercises)
}
