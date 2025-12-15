package services

import (
	"teomebot/config"
	"teomebot/repositories"
	"testing"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupPointsServiceTest() *gorm.DB {

	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	db.AutoMigrate(
		&repositories.TwitchUser{},
		&repositories.PresentUser{},
		&repositories.StreakPresentUser{},
	)

	user1 := repositories.TwitchUser{
		UUID:       "user-uuid-1",
		TwitchId:   "twitch-id-1",
		TwitchNick: "testuser1",
	}

	user2 := repositories.TwitchUser{
		UUID:       "user-uuid-2",
		TwitchId:   "twitch-id-2",
		TwitchNick: "testuser2",
	}

	db.Create(&user1)
	db.Create(&user2)

	present1 := repositories.PresentUser{
		UUID:      "123456789",
		UserID:    user1.UUID,
		CreatedAt: time.Now().AddDate(0, 0, -1),
	}

	present2 := repositories.PresentUser{
		UUID:      "1234567890",
		UserID:    user1.UUID,
		CreatedAt: time.Now(),
	}

	db.Create(&present1)
	db.Create(&present2)

	return db
}

func TestExistsPresent(t *testing.T) {

	db := SetupPointsServiceTest()

	settings := &config.Config{
		LoyaltyServiceURI: "http://localhost:8081",
	}

	service := NewPointsService(settings, db)

	user1 := twitch.User{
		ID:          "twitch-id-1",
		DisplayName: "testuser1",
	}
	msg1, _ := service.MgmtPresenca(user1)
	assert.Equal(t, "testuser1 você já assinou presença hoje!", msg1)

}
