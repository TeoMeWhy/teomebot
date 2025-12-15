package repositories

import (
	"fmt"

	"gorm.io/gorm"
)

type TwitchUser struct {
	UUID       string `json:"uuid" gorm:"primaryKey"`
	TwitchId   string `json:"twitch" gorm:"unique"`
	TwitchNick string `json:"twitch_nick"`
}

type UserRepository struct {
	ConDB *gorm.DB
}

func (r *UserRepository) CreateUser(t *TwitchUser) error {
	return r.ConDB.Create(t).Error
}

func (r *UserRepository) Update(t *TwitchUser) error {
	return r.ConDB.Save(t).Error
}

func (r *UserRepository) GetUserByField(fieldName, fieldValue string) (*TwitchUser, error) {

	user := &TwitchUser{}
	res := r.ConDB.First(&user, fmt.Sprintf("%s = ?", fieldName), fieldValue)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{ConDB: db}
}
