package model

import "github.com/google/uuid"

type Exercise struct {
	ID   uuid.UUID `gorm:"primaryKey"`
	Name string
}

func GetExercises() ([]Exercise, error) {
	var exercises []Exercise

	err := db.Find(&exercises).Error

	return exercises, err
}
