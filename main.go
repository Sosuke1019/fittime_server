package main

import (
	"github.com/labstack/echo/v4"
	"github.com/ponyo-E/fittime_server/model"
	"github.com/ponyo-E/fittime_server/router"
)

func main() {
	sqlDB := model.DBConnection()
	defer sqlDB.Close()
	e := echo.New()
	router.SetRouter(e)
}
