package model

import (
	"errors"

	"github.com/google/uuid"
)

// User型を定義する
type User struct {
	ID       uuid.UUID
	Name     string
	Mail     string
	Password []byte
	Profile  string
	Path     string
	Point    int
}

func CreateUser(username string, mail string, password string) error {
	var count int64
	db.Model(&User{}).Where("mail = ?", mail).Count(&count)
	if count != 0 {
		return errors.New("mail already exists")
	}
	id, err := uuid.NewUUID()

	newUser := User{
		ID:       id,
		Name:     username,
		Mail:     mail,
		Password: pass,
		Profile:  "",
		Path:     "",
		Point:    0,
	}

	err = db.Create(&newUser).Error
	return err
}
