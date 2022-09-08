package model

import (
	"fmt"

	"github.com/google/uuid"
)

type Menu struct {
	ID         uuid.UUID `gorm:"primaryKey"`
	ExerciseID uuid.UUID `gorm:"size:30"`
	Exercise   Exercise
	MenuID     uuid.UUID `gorm:"size:40"`
	Menu       Menu
	No         int
	Time       int
}

type Menu struct {
	ID            uuid.UUID `gorm:"primaryKey"`
	Title         string
	UserID        uuid.UUID `gorm:"size:30"`
	User          User
	Body          string
	Path          string
	Nice          int
	Point         int
	ExerciseParts []ExercisePart
}

func SearchMenu(word string) ([]Menu, error) {
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
