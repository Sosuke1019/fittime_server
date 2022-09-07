package router

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ponyo-E/fittime_server/model"
	"net/http"
)

type ReqExercisePart struct {
	ExerciseId uuid.UUID `json:"ExerciseId"`
	Time       int       `json:"time"`
}
type ReqMenu struct {
	Title         string            `json:"title"`
	Body          string            `json:"body"`
	ExerciseParts []ReqExercisePart `json:"exercises"`
}

func PostMenuHandler(c echo.Context) error {

	userId, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	var req ReqMenu
	err = c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	id, _ := uuid.NewUUID()

	var exerciseParts []model.ExercisePart
	for i, r := range req.ExerciseParts {
		ExercisePartId, _ := uuid.NewUUID()
		exercisePart := model.ExercisePart{
			ID:         ExercisePartId,
			ExerciseID: r.ExerciseId,
			MenuID:     id,
			No:         i,
			Time:       r.Time,
		}
		exerciseParts = append(exerciseParts, exercisePart)
	}

	// ReqMenu型からMenu型に
	menu := model.Menu{
		ID:            id,
		Title:         req.Title,
		UserID:        userId,
		Body:          req.Body,
		Path:          "",
		Nice:          0,
		Point:         0,
		ExerciseParts: exerciseParts,
	}

	err = model.AddMenu(menu)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	return c.NoContent(http.StatusOK)
}
