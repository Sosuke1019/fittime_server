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

func AddFavorite(favorite Favorite) error {
	err := db.Create(&favorite).Error
	if err != nil {
		return err
	}
	return nil
}

func GetFavorite(userId uuid.UUID) ([]*ResMenu, error) {
	var menus []Menu
	err := db.Table("menus").Select("menus.id,menus.title,menus.user_id,users.name,menus.body,menus.nice,menus.point,menus.exercise_id,exercises.name").
		Joins("left join users on users.id = menus.user_id").
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
			MenuId:   menu.ID,
			Title:    menu.Title,
			UserId:   menu.UserID,
			UserName: menu.User.Name,
			Body:     menu.Body,
			Nice:     menu.Nice,
			Point:    menu.Point,
			Exercises: ResExercise{
				ExerciseId: menu.ExerciseID,
				Name:       menu.Exercise.Name,
			},
		})
	}
	fmt.Println(resMenu)
	return resMenu, nil
}
