package model

import (
	"golang.org/x/crypto/bcrypt"
)

// パスワードをハッシュ化する
func HashPassword(pass string) []byte {
	convert := []byte(pass)
	hash, _ := bcrypt.GenerateFromPassword(convert, 8)
	return hash
}

func CheckPassword(mail string, pass string) (*User, error) {

	var user User
	// mailからpasswordを取得
	err := db.Model(&User{}).Where("mail = ?", mail).Find(&user).Error
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(pass))
	if err != nil {
		return nil, err
	}

	return &user, nil

}
