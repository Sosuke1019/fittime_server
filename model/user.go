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

func GetLevelAndStatus(point int) (int, string) {
	level := point / 10
	status := "見習い勇者"
	return level, status
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

func GetUser(userId uuid.UUID) (User, error) {
	var user User
	err := db.Model(&User{}).Where("id = ?", userId).Find(&user).Error

	return user, err
}

func UpdateProfile(userId uuid.UUID, profile string) error {
	err := db.Model(&User{}).Where("id = ?", userId).Update("Profile", profile).Error

	return err
}

func UpdateName(userId uuid.UUID, name string) error {
	err := db.Model(&User{}).Where("id = ?", userId).Update("Name", name).Error

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
