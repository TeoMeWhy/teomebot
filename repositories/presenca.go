package repositories

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PresentUser struct {
	UUID      string    `json:"uuid" gorm:"primaryKey"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type StreakPresentUser struct {
	UUID      string    `json:"uuid" gorm:"primaryKey"`
	UserID    string    `json:"user_id"`
	Qtd       int64     `json:"qtd"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
}

func NewStreakPresentUser(twitch_id string) *StreakPresentUser {
	return &StreakPresentUser{
		UUID:   uuid.New().String(),
		UserID: twitch_id,
		Qtd:    1,
	}
}

type PresencaRepository struct {
	db *gorm.DB
}

func NewPresencaRepository(db *gorm.DB) *PresencaRepository {
	return &PresencaRepository{
		db: db,
	}
}

func (r *PresencaRepository) CreateStreak(twitch_id string) error {
	streak := NewStreakPresentUser(twitch_id)
	return r.db.Create(&streak).Error
}

func (r *PresencaRepository) UpdateStreak(streak *StreakPresentUser) error {
	return r.db.Save(&streak).Error
}

func (r *PresencaRepository) LoadLastUpdatedStreak(twitch_id string) (*StreakPresentUser, error) {
	streak := &StreakPresentUser{}
	res := r.db.Where("user_id = ?", twitch_id).Order("updated_at DESC").First(&streak)
	return streak, res.Error
}
