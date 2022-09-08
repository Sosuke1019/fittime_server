package router

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ponyo-E/fittime_server/model"
)

func GetFavoriteHandler(c echo.Context) error {
	userId, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	resMenu, err := model.GetFavorite(userId)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	return c.JSON(http.StatusOK, resMenu)
}

func PostFavoriteHandler(c echo.Context) error {
	userId, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	menuId, err := uuid.Parse(c.Param("menuId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	favorite := model.Favorite{UserID: userId, MenuID: menuId}
	err = model.AddFavorite(favorite)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	return c.NoContent(http.StatusOK)
}
