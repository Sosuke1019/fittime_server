package model

import (
	"time"

	"github.com/google/uuid"
)

type Log struct {
	ID     uuid.UUID `gorm:"primaryKey"`
	UserID uuid.UUID `gorm:"size:40"`
	User   User
	Date   time.Time
	MenuId uuid.UUID `gorm:"size:40"`
	Menu   Menu
}

func GetTimeline() ([]Menu, error) {
	var menu []Menu
	err := db.Table("logs").Select("menus.id, menus.title").Joins("left join menus on menus.id = logs.menu_id").
		Limit(5).Find(&menu).Error
	if err != nil {
		return nil, err
	}
	return menu, nil
}

// logデータベースから降順 -> 多分○
// limitで上から5つに制限する -> ○
// ログのmenu.idからmenuの名前を持ってくる
func AddLog(userId uuid.UUID, menuId uuid.UUID) error {
	id, err := uuid.NewUUID()

	date := time.Now()

	newLog := Log{
		ID:     id,
		UserID: userId,
		Date:   date,
		MenuId: menuId,
	}

	err = db.Model(&Log{}).Create(&newLog).Error
	return err
}
