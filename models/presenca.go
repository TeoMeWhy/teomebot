package models

import (
	"log"
	"time"

	"github.com/google/uuid"
)

type PresentUser struct {
	UUID      string    `json:"uuid" gorm:"primaryKey"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func PresentExists(userId string) bool {
	p := &PresentUser{}
	res := conDB.Where("user_id = ?", userId).Order("created_at desc").First(&p)
	if res.Error != nil {
		return false
	}

	dtCreated := p.CreatedAt.Format("2006-01-02")
	dtNow := time.Now().Format("2006-01-02")
	return dtNow == dtCreated
}

func NewPresentUser(userID string) *PresentUser {
	return &PresentUser{
		UUID:   uuid.New().String(),
		UserID: userID,
	}
}

type StreakPresentUser struct {
	UUID      string    `json:"uuid" gorm:"primaryKey"`
	UserID    string    `json:"user_id"`
	Qtd       int64     `json:"qtd"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
}

func NewStreakPresentUser(userID string) *StreakPresentUser {
	return &StreakPresentUser{
		UUID:   uuid.New().String(),
		UserID: userID,
		Qtd:    0,
	}
}

func LoadStreak(userID string) *StreakPresentUser {

	streak := &StreakPresentUser{}
	res := conDB.Where("user_id = ?", userID).Order("updated_at DESC").First(&streak)
	if res.Error != nil {
		log.Println(res)
		return NewStreakPresentUser(userID)
	}

	now := time.Now().Year() + time.Now().YearDay()
	updated := streak.UpdatedAt.Year() + streak.UpdatedAt.YearDay()

	if now-updated <= 1 {
		return streak
	}

	return NewStreakPresentUser(userID)
}
