package models

import (
	"fmt"

	"gorm.io/gorm"
)

// ConDB.AutoMigrate(&models.TwitchUser{})

type TwitchUser struct {
	UUID       string `json:"uuid" gorm:"primaryKey"`
	TwitchId   string `json:"twitch" gorm:"unique"`
	TwitchNick string `json:"twitch_nick"`
}

func (tu *TwitchUser) Create(con *gorm.DB) error {

	res := con.Create(tu)
	defer res.Commit()
	if res.Error != nil {
		res.Rollback()
		return res.Error
	}

	return nil
}

func (tu *TwitchUser) Update(con *gorm.DB) error {
	res := conDB.Save(tu)
	defer res.Commit()
	if res.Error != nil {
		res.Rollback()
		return res.Error
	}

	return nil
}

func GetUserByField(fieldName, fieldValue string, con *gorm.DB) (*TwitchUser, error) {

	user := &TwitchUser{}

	res := con.First(&user, fmt.Sprintf("%s = ?", fieldName), fieldValue)
	defer res.Commit()
	if res.Error != nil {
		res.Rollback()
		return nil, res.Error
	}
	return user, nil
}
