package router

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ponyo-E/fittime_server/model"
	"net/http"
)

type ReqMenu struct {
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	ExerciseId uuid.UUID `json:"exerciseId"`
	Time       int       `json:"time"`
}

func PostMenuHandler(c echo.Context) error {

	userId, err := uuid.Parse(c.Param("userId"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, c.Param("userId"))
	}

	var req ReqMenu
	err = c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bind Request")
	}

	menuId, _ := uuid.NewUUID()

	// ReqMenu型からMenu型に
	menu := model.Menu{
		ID:         menuId,
		Title:      req.Title,
		UserID:     userId,
		Body:       req.Body,
		Path:       "",
		Nice:       0,
		Point:      req.Time,
		ExerciseID: req.ExerciseId,
	}

	err = model.AddMenu(menu)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.NoContent(http.StatusOK)
}
