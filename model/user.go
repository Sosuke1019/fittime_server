package model

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// User型を定義する
type User struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Name     string
	Mail     string
	Password []byte `gorm:"type:VARCHAR(200)"`
	Profile  string
	Path     string
	Point    int
}

func CreateUser(username string, mail string, pass string) error {
	var count int64
	db.Model(&User{}).Where("mail = ?", mail).Count(&count)
	if count != 0 {
		return errors.New("mail already exists")
	}
	id, _ := uuid.NewUUID()

	hash := HashPassword(pass)

	fmt.Println(hash)

	newUser := User{
		ID:       id,
		Name:     username,
		Mail:     mail,
		Password: hash,
		Profile:  "",
		Path:     "",
		Point:    0,
	}

	err := db.Create(&newUser).Error
	return err
}

func AddProfile(userId uuid.UUID, profile string) error {
	err := db.Model(&User{}).Where("id = ?", userId).Update("Profile", profile).Error

	return err
}

func SearchUser(word string) ([]User, error) {
	var users []User
	err := db.Where("name LIKE ?", "%"+word+"%").Find(&users).Error
	if err != nil {
		return nil, err
	}

	if len(users) > 5 {
		users = users[:5]
	}

	return users, nil
}
