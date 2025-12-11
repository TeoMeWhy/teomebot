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

func NewPresentUser(twitchUser *TwitchUser) *PresentUser {
	return &PresentUser{
		UUID:   uuid.New().String(),
		UserID: twitchUser.UUID,
	}
}

type StreakPresentUser struct {
	UUID      string    `json:"uuid" gorm:"primaryKey"`
	UserID    string    `json:"user_id"`
	Qtd       int64     `json:"qtd"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
}

func NewStreakPresentUser(twitchUser *TwitchUser) *StreakPresentUser {
	return &StreakPresentUser{
		UUID:   uuid.New().String(),
		UserID: twitchUser.UUID,
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

func (r *PresencaRepository) CreatePresenca(twitchUser *TwitchUser) (*PresentUser, error) {
	present := NewPresentUser(twitchUser)
	return present, r.db.Create(&present).Error
}

func (r *PresencaRepository) DeletePresenca(present *PresentUser) error {
	return r.db.Delete(&present).Error
}

func (r *PresencaRepository) LoadLastPresent(twitchUser *TwitchUser) (*PresentUser, error) {
	present := &PresentUser{}
	res := r.db.Where("user_id = ?", twitchUser.UUID).Order("created_at DESC").First(&present)
	return present, res.Error
}

func (r *PresencaRepository) CreateStreak(twitchUser *TwitchUser) error {
	streak := NewStreakPresentUser(twitchUser)
	return r.db.Create(&streak).Error
}

func (r *PresencaRepository) UpdateStreak(streak *StreakPresentUser) error {
	return r.db.Save(&streak).Error
}

func (r *PresencaRepository) LoadLastUpdatedStreak(twitchUser *TwitchUser) (*StreakPresentUser, error) {
	streak := &StreakPresentUser{}
	res := r.db.Where("user_id = ?", twitchUser.UUID).Order("updated_at DESC").First(&streak)
	return streak, res.Error
}
