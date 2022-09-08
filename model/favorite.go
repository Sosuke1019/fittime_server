package model

import "github.com/google/uuid"

type Favorite struct {
	ID     uuid.UUID `gorm:"primaryKey"`
	UserID uuid.UUID `gorm:"size:40"`
	User   User
	MenuID uuid.UUID `gorm:"size:40"`
	Menu   Menu
}

func AddFavorite(favorite Favorite) error {
	err := db.Create(&favorite).Error
	if err != nil {
		return err
	}
	return nil
}

// func GetFavorite(userId uuid.UUID) error {
// 	err := db.Table('')
// }
