package utils

import (
	"teomebot/config"
	"teomebot/repositories"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDBConnection(settings *config.Config) (*gorm.DB, error) {

	// godotenv.Load(".env")
	// HOST_DB := os.Getenv("HOST_DB")
	// PORT_DB := os.Getenv("PORT_DB")
	// USER_DB := os.Getenv("USER_DB")
	// PASSWORD_DB := os.Getenv("PASSWORD_DB")
	// dsn := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn = fmt.Sprintf(dsn, USER_DB, PASSWORD_DB, HOST_DB, PORT_DB, "teomebot")

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
