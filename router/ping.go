package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetPingHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
