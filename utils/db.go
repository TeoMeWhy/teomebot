package utils

import (
	"teomebot/config"
	"teomebot/repositories"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDBConnection(settings *config.Config) (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(settings.DsnMysql), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&repositories.PresentUser{},
		&repositories.StreakPresentUser{},
		&repositories.TwitchUser{},
		&repositories.Messagem{},
	)

	return db, err
}
