package model

import (
	"github.com/google/uuid"
)

// User型を定義する
type User struct {
	ID       uuid.UUID `gorm:"primary_key"`
	name     string
	mail     string
	password string
	profile  string
	path     string
	point    int
}

func CreateUser(username string, mail string, password string) error {
	var count int64
	db.Model(&User{}).Where("mail = ?", mail).Count(&count)
	if count != 0 {
		return "Bad Request"
	}
	id, err := uuid.NewUUID()
	newUser := User{
		ID:       id,
		name:     username,
		mail:     mail,
		password: password,
		profile:  "",
		path:     "",
		point:    0,
	}

	err = db.Create(&newUser).Error
	return err
}
