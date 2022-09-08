package model

import (
	"strconv"

	"github.com/google/uuid"
)

var exerciseList = []string{"腹筋", "スクワット", "腕立てふせ"}

func CreateExercise() {
	for _, l := range exerciseList {
		id, _ := uuid.NewUUID()
		exercise := Exercise{
			ID:   id,
			Name: l,
		}
		db.Create(&exercise)
	}
}

func CreateDummy() {
	for i := 0; i < 10; i++ {
		CreateUser("username"+strconv.Itoa(i), strconv.Itoa(i)+"@test.com",
			"password"+strconv.Itoa(i))
	}
}
