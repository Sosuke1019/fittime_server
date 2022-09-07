package model

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// パスワードをハッシュ化する
func HashPassword(password string) ([]byte, error) {
	convert := []byte(password)
	hashedPass, err := bcrypt.GenerateFromPassword(convert, bcrypt.DefaultCost)
	return hashedPass, err
}

func CheckPassword(mail string, password string) (uuid.UUID, error) {
	hashedPass, err := HashPassword(password)
	if err != nil {
		return "", err
	}

	var user User
	// mailからpasswordを取得
	err = db.Model(&User{}).Where("mail = ?", mail).Find(&user).Error
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword(hashedPass, user.Password)
	if err != nil {
		return "", err
	}

	return user.ID, nil

}
