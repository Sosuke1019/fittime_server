package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Log struct {
	ID     uuid.UUID `gorm:"primaryKey"`
	UserID uuid.UUID `gorm:"size:40"`
	User   User
	Date   time.Time
	MenuId uuid.UUID `gorm:"size:40"`
	Menu   Menu
}

type Timeline struct {
	User
	Date time.Time
	Menu
	Exercise
}

type ResTimeline struct {
	UserId   uuid.UUID `json:"userId"`
	UserName string    `json:"userName"`
	Menu     ResMenu   `json:"menu"`
	Date     time.Time `json:"date"`
}

func GetTimeline() ([]ResTimeline, error) {
	var logs []Timeline
	err := db.Table("logs").Select("users.*,logs.date,menus.*,exercises.*").
		Joins("left join menus on menus.id = logs.menu_id").
		Joins("left join users on users.id = logs.user_id").
		Joins("left join exercises on exercises.id = menus.exercise_id").
		Order("logs.date").Limit(5).Find(&logs).Error
	if err != nil {
		return nil, err
	}
	var resTimeline []ResTimeline
	for _, data := range logs {
		fmt.Println(data)
		resTimeline = append(resTimeline, ResTimeline{
			UserId:   data.UserID,
			UserName: data.User.Name,
			Menu: ResMenu{
				MenuId:   data.Menu.ID,
				Title:    data.Menu.Title,
				UserId:   data.UserID,
				UserName: data.User.Name,
				Body:     data.Menu.Body,
				Nice:     data.Menu.Nice,
				Point:    data.Menu.Nice,
				Exercises: ResExercise{
					ExerciseId: data.Menu.ExerciseID,
					Name:       data.Exercise.Name,
				},
			},
			Date: data.Date,
		})
	}
	return resTimeline, nil
}

// logデータベースから降順 -> 多分○
// limitで上から5つに制限する -> ○
// ログのmenu.idからmenuの名前を持ってくる
func AddLog(userId uuid.UUID, menuId uuid.UUID) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	date := time.Now()

	newLog := Log{
		ID:     id,
		UserID: userId,
		Date:   date,
		MenuId: menuId,
	}

	err = db.Model(&Log{}).Create(&newLog).Error
	return err
}

func GetMyTimeline(userId uuid.UUID) ([]ResTimeline, error) {
	var logs []Timeline
	err := db.Table("logs").Select("users.*,logs.date,menus.*,exercises.*").
		Joins("left join menus on menus.id = logs.menu_id").
		Joins("left join users on users.id = logs.user_id").
		Joins("left join exercises on exercises.id = menus.exercise_id").
		Where("users.id = ?", userId).
		Order("logs.date").Limit(5).Find(&logs).Error
	if err != nil {
		return nil, err
	}
	var resTimeline []ResTimeline
	for _, data := range logs {
		fmt.Println(data)
		resTimeline = append(resTimeline, ResTimeline{
			UserId:   data.UserID,
			UserName: data.User.Name,
			Menu: ResMenu{
				MenuId:   data.Menu.ID,
				Title:    data.Menu.Title,
				UserId:   data.UserID,
				UserName: data.User.Name,
				Body:     data.Menu.Body,
				Nice:     data.Menu.Nice,
				Point:    data.Menu.Nice,
				Exercises: ResExercise{
					ExerciseId: data.Menu.ExerciseID,
					Name:       data.Exercise.Name,
				},
			},
			Date: data.Date,
		})
	}
	return resTimeline, nil
}
