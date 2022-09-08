package model

import (
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
