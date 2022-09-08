package model

import (
	"strconv"

	"github.com/google/uuid"
)

var exerciseList = []string{"腹筋", "スクワット", "腕立てふせ"}

func CreateExercise() []uuid.UUID {
	var exerciseIdList []uuid.UUID
	for _, l := range exerciseList {
		id, _ := uuid.NewUUID()
		exercise := Exercise{
			ID:   id,
			Name: l,
		}
		exerciseIdList = append(exerciseIdList, id)
		db.Create(&exercise)
	}
	return exerciseIdList
}

func CreateUserDummy() {
	for i := 0; i < 10; i++ {
		CreateUser("username"+strconv.Itoa(i), strconv.Itoa(i)+"@test.com",
			"password"+strconv.Itoa(i))
	}
}

func CreateDummy() {
	CreateUserDummy()
	exerciseIdList := CreateExercise()
	var users []User
	db.Find(&users)
	for _, user := range users {
		for _, exerciseId := range exerciseIdList {
			id, _ := uuid.NewUUID()
			menu := Menu{ID: id, UserID: user.ID, ExerciseID: exerciseId}
			db.Create(&menu)
			id, _ = uuid.NewUUID()
			favorite := Favorite{ID: id, MenuID: menu.ID, UserID: user.ID}
			db.Create(&favorite)
		}
	}
}
