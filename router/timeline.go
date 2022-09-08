package router

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ResTimeline struct {
	userId   uuid.UUID `json:"userId"`
	Username string    `json:"username"`
	Menu     string    `json:"menu"`
	Date     time.Time `json:"time"`
}

func TimelineHandler(c echo.Context) error {
	return c.JSON(http.StatusOK)
}

//ログから最新の4つ持ってくる
//ログのmenu.idからmenuの名前を持ってくる
//データを整形して変えす
