package model

import (
	"github.com/google/uuid"
	"time"
)

type Log struct {
	ID     uuid.UUID `gorm:"primaryKey"`
	UserID uuid.UUID `gorm:"size:30"`
	User   User
	Date   time.Time
	MenuId uuid.UUID `gorm:"size:30"`
	Menu   Menu
}
