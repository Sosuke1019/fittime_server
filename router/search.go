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

type ResExercisePart struct {
	no   int    `json:"no"`
	Name string `json:"name"`
	time int    `json:"time"`
}

type ResMenu struct {
	MenuId    uuid.UUID         `json:"menuId"`
	Title     string            `json:"title"`
	UserName  string            `json:"username"`
	Body      string            `json:"body"`
	Nice      int               `json:"nice"`
	Point     int               `json:"point"`
	Exercises []ResExercisePart `json:"exercises"`
}
type ResSearch struct {
	ResUser []ResUser `json:"users"`
	ResMenu []ResMenu `json:"menus"`
}

func SearchHandler(c echo.Context) error {
	word := c.QueryParam("word")

	users, err := model.SearchUser(word)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	// User型をResUser型に変換
	var resUsers []ResUser
	for _, user := range users {
		resUser := ResUser{
			UserId: user.ID,
			Name:   user.Name,
		}
		resUsers = append(resUsers, resUser)
	}

	menus, err := model.SearchMenu(word)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	// Menu型をResMenu型に変換
	var resMenus []ResMenu
	for _, menu := range menus {
		// ExercisePart型をResExercisePart型に変換
		var resExerciseParts []ResExercisePart
		for _, exercisePart := range menu.ExerciseParts {
			resExercisePart := ResExercisePart{
				no:   exercisePart.No,
				Name: exercisePart.Exercise.Name,
				time: exercisePart.Time,
			}
			resExerciseParts = append(resExerciseParts, resExercisePart)
		}

		resMenu := ResMenu{
			MenuId:    menu.ID,
			Title:     menu.Title,
			UserName:  menu.User.Name,
			Body:      menu.Body,
			Nice:      menu.Nice,
			Point:     menu.Point,
			Exercises: resExerciseParts,
		}
		resMenus = append(resMenus, resMenu)
	}

	resSearch := ResSearch{
		ResUser: resUsers,
		ResMenu: resMenus,
	}

	return c.JSON(http.StatusOK, resSearch)

}
