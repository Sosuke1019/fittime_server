package model

import (
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
	CreateUser("Taro", "taro@gmail.com", "password")
	CreateUser("Yumi", "yumi@icloud.com", "password")
	CreateUser("Kenta", "kenta@gmail.com", "password")
	CreateUser("Mei", "Mei@icloud.com", "password")
	CreateUser("Ryosuke", "ryosuke@icloud.com", "password")
}

func CreateDummy() {
	CreateUserDummy()
	exerciseIdList := CreateExercise()
	var users []User
	db.Find(&users)
		exerciseIdList {
			id, _ := uuid.NewUUID()
			menu := Menu{
				ID:         id,
				Title:      "ムキムキメニュー",
				UserID:     user.ID,
				Body:       "ムキムキになります",
				Path:       "",
				Nice:       2,
				Point:      10,
				ExerciseID: exerciseId,
			}
			db.Create(&menu)
			id, _ = uuid.NewUUID()
			favorite := Favorite{ID: id, MenuID: menu.ID, UserID: user.ID}
			db.Create(&favorite)
		}
		exerciseIdList {
			id, _ := uuid.NewUUID()
			title := "menuTitle"
			body := "menuBody"
			menu := Menu{
				ID:         id,
				Title:      "朝の運動",
				UserID:     user.ID,
				Body:       "程よい汗をかけます",
				Path:       "",
				Nice:       4,
				Point:      3,
				ExerciseID: exerciseId,
			}
			db.Create(&menu)
			id, _ = uuid.NewUUID()
			favorite := Favorite{ID: id, MenuID: menu.ID, UserID: user.ID}
			db.Create(&favorite)
		}
		exerciseIdList {
			id, _ := uuid.NewUUID()
			title := "menuTitle"
			body := "menuBody"
			menu := Menu{
				ID:         id,
				Title:      "家族で運動",
				UserID:     user.ID,
				Body:       "家族みんなで汗をかきましょう",
				Path:       "",
				Nice:       80,
				Point:      7,
				ExerciseID: exerciseId,
			}
			db.Create(&menu)
			id, _ = uuid.NewUUID()
			favorite := Favorite{ID: id, MenuID: menu.ID, UserID: user.ID}
			db.Create(&favorite)
		}
		exerciseIdList {
			id, _ := uuid.NewUUID()
			title := "menuTitle"
			body := "menuBody"
			menu := Menu{
				ID:         id,
				Title:      "激しいメニュー",
				UserID:     user.ID,
				Body:       "限界突破しましょう！",
				Path:       "",
				Nice:       6,
				Point:      10,
				ExerciseID: exerciseId,
			}
			db.Create(&menu)
			id, _ = uuid.NewUUID()
			favorite := Favorite{ID: id, MenuID: menu.ID, UserID: user.ID}
			db.Create(&favorite)
		}
		exerciseIdList {
			id, _ := uuid.NewUUID()
			title := "menuTitle"
			body := "menuBody"
			menu := Menu{
				ID:         id,
				Title:      "休日ダイエット",
				UserID:     user.ID,
				Body:       "休日もダイエットを継続!",
				Path:       "",
				Nice:       20,
				Point:      6,
				ExerciseID: exerciseId,
			}
			db.Create(&menu)
			id, _ = uuid.NewUUID()
			favorite := Favorite{ID: id, MenuID: menu.ID, UserID: user.ID}
			db.Create(&favorite)
		}
	}
}
