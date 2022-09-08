package model

import (
	"fmt"

	"github.com/google/uuid"
)

type Menu struct {
	ID         uuid.UUID `gorm:"primaryKey"`
	Title      string
	UserID     uuid.UUID `gorm:"size:40"`
	User       User      `gorm:"foreignKey:UserID"`
	Body       string
	Path       string
	Nice       int
	Point      int
	ExerciseID uuid.UUID `gorm:"size:40"`
	Exercise   Exercise
}

func AddMenu(menu Menu) error {
	err := db.Create(&menu).Error
	if err != nil {
		return err
	}

	return nil
}

type ResExercise struct {
	ExerciseId uuid.UUID `json:"exerciseId"`
	Name       string    `json:"exerciseName"`
}

type ResMenu struct {
	MenuId    uuid.UUID   `json:"menuId"`
	Title     string      `json:"title"`
	UserId    uuid.UUID   `json:"userId"`
	UserName  string      `json:"username"`
	Body      string      `json:"body"`
	Nice      int         `json:"nice"`
	Point     int         `json:"point"`
	Exercises ResExercise `json:"exercises"`
}
type ResSearch struct {
	ResMenus []ResMenu `json:"menus"`
}

func SearchMenu(word string) (*ResSearch, error) {
	var menus []Menu
	err := db.Where("body LIKE ?", "%"+word+"%").Find(&menus).Error
	if err != nil {
		return nil, err
	}

	if len(menus) > 5 {
		menus = menus[:5]
	}

	var user User
	var exercise Exercise
	var resMenus []ResMenu
	for _, menu := range menus {

		//get username
		err := db.Where("id = ?", menu.UserID).Find(&user).Error
		if err != nil {
			return nil, err
		}

		// get exercise
		err = db.Where("id = ?", menu.ExerciseID).Find(&exercise).Error
		if err != nil {
			return nil, err
		}

		// exercise → ResExercise
		resExercise := ResExercise{
			ExerciseId: exercise.ID,
			Name:       exercise.Name,
		}

		fmt.Println(resExercise)

		// menu → ResMenu
		resMenu := ResMenu{
			MenuId:    menu.ID,
			Title:     menu.Title,
			UserId:    menu.UserID,
			UserName:  user.Name,
			Body:      menu.Body,
			Nice:      menu.Nice,
			Point:     menu.Point,
			Exercises: resExercise,
		}

		fmt.Println(resMenu)

		resMenus = append(resMenus, resMenu)

	}

	// ResMenu → ResSearch
	resSearch := ResSearch{
		ResMenus: resMenus,
	}

	return &resSearch, nil
}
