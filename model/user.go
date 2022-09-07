package model

import (
	"errors"

	"github.com/google/uuid"
)

// User型を定義する
type User struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Name     string
	Mail     string
	Password string
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
		Password: password,
		Profile:  "",
		Path:     "",
		Point:    0,
	}

	err = db.Create(&newUser).Error
	return err
}

func AddProfile(profile string) error {
	db.Model(&User{}).Where("id = ?", true).Update(&profile)
}


ユーザーidはパラメータで取得
新しいprofileを入れる