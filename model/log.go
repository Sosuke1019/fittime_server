package model

import (
	"time"

	"github.com/google/uuid"
)

type Log struct {
	ID     uuid.UUID `gorm:"primaryKey"`
	UserID uuid.UUID `gorm:"size:40"`
	Date   time.Time
	MenuId uuid.UUID `gorm:"size:40"`
}

func AddLog(userId uuid.UUID, menuId uuid.UUID) error {
	id, err := uuid.NewUUID()

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
