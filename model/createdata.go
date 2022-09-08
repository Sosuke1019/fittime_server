package model

import "github.com/google/uuid"

var exerciseList = []string{"腹筋", "スクワット", "腕立てふせ"}

func CreateData() {
	for _, l := range exerciseList {
		id, _ := uuid.NewUUID()
		exercise := Exercise{
			ID:   id,
			Name: l,
		}
		db.Create(&exercise)
	}

}
