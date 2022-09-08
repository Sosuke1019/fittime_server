package router

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ponyo-E/fittime_server/model"
)

type ResTimeline struct {
	UserId   uuid.UUID `json:"userId"`
	Username string    `json:"username"`
	Menu     string    `json:"menu"`
	Date     time.Time `json:"time"`
}

func GetTimelineHandler(c echo.Context) error {
	// DBからTimeLineのデータを取得

	logs, err := model.GetTimeline()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if len(logs) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "GetTimeline取得エラー")
	}

	// DBからmenuを取得
	//logs[0].MenuId こんな感じでMenuIdがとれます
	// 参考になるのはModel>user.go>GetUser()

	// MenuIdをMenuに変えて必要なデータを返す

	return c.JSON(http.StatusOK, logs)
}

func GetMyTimelineHandler(c echo.Context) error {
	userId, err := uuid.Parse(c.Param("userId"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "uuidエラー")
	}

	logs, err := model.GetMyTimeline(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, logs)
}
