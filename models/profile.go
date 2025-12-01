package models

import (
	"time"

	"github.com/google/uuid"
)

type ProfileUser struct {
	UUID      string    `json:"uuid" gorm:"primaryKey"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func ProfileUserExists(userId string) bool {
	p := &ProfileUser{}
	res := conDB.Where("user_id = ?", userId).Order("created_at desc").First(&p)
	if res.Error != nil {
		return false
	}

	dtCreated := p.CreatedAt.Format("2006-01-02")
	dtNow := time.Now().Format("2006-01-02")
	return dtNow == dtCreated
}

func NewProfileUser(userID string) *ProfileUser {
	return &ProfileUser{
		UUID:   uuid.New().String(),
		UserID: userID,
	}
}
