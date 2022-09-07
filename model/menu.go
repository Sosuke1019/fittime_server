package model

import "github.com/google/uuid"

type Exsercise struct {
	ID   uuid.UUID `gorm:"primaryKey"`
	Name string
	Time int
}

type Menu struct {
	ID     uuid.UUID `gorm:"primaryKey"`
	UserID uuid.UUID `gorm:"size:30"`
	User   User
	Body   string
	Path   string
	Nice   int
	Point  int
}
