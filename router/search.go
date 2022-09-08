package router

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ponyo-E/fittime_server/model"
	"net/http"
)

type ResUser struct {
	UserId uuid.UUID `json:"userId"`
	Name   string    `json:"username"`
}

func SearchHandler(c echo.Context) error {
	word := c.QueryParam("word")

	resSearch, err := model.SearchMenu(word)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, resSearch)

}
