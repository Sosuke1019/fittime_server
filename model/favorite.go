package model

import (
	"fmt"

	"github.com/google/uuid"
)

type Favorite struct {
	ID     uuid.UUID `gorm:"primaryKey"`
	UserID uuid.UUID `gorm:"size:40"`
	User   User
	MenuID uuid.UUID `gorm:"size:40"`
	Menu   Menu
}

type FavoriteData struct {
	User
	Menu
	Exercise
}

func AddFavorite(favorite Favorite) error {
	err := db.Create(&favorite).Error
	if err != nil {
		return err
	}
	return nil
}

func GetFavorite(userId uuid.UUID) ([]*ResMenu, error) {
	var menus []FavoriteData
	err := db.Table("favorites").Select("users.*,menus.*,exercises.*").
		Joins("left join menus on menus.id = favorites.menu_id").
		Joins("left join users on users.id = favorites.user_id").
		Joins("left join exercises on exercises.id = menus.exercise_id").
		Where("users.id = ?", userId).
		Limit(5).
		Find(&menus).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(menus)
	var resMenu []*ResMenu
	for _, menu := range menus {
		fmt.Println(menu)
		resMenu = append(resMenu, &ResMenu{
			MenuId:   menu.Menu.ID,
			Title:    menu.Menu.Title,
			UserId:   menu.User.ID,
			UserName: menu.User.Name,
			Body:     menu.Menu.Body,
			Nice:     menu.Menu.Nice,
			Point:    menu.Menu.Point,
			Exercises: ResExercise{
				ExerciseId: menu.Exercise.ID,
				Name:       menu.Exercise.Name,
			},
		})
	}
	fmt.Println(resMenu)
	return resMenu, nil
}
