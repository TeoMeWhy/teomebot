package services

import (
	"teomebot/config"
	"teomebot/errors"
	"teomebot/repositories"
	"testing"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewUser(t *testing.T) {

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&repositories.TwitchUser{})

	settings := config.Config{
		LoyaltyServiceURI: "http://localhost:8081",
	}

	perfilService := NewPerfilService(&settings, db)

	twitchUser := twitch.User{
		ID:          "1234",
		DisplayName: "My_Name",
		Name:        "my_name",
	}

	msg, err := perfilService.CreateNewUser(twitchUser)
	assert.Equal(t, "My_Name usuário criado com sucesso", msg)
	assert.NoError(t, err)

	user, _ := perfilService.userRepository.GetUserByField("twitch_id", twitchUser.ID)
	perfilService.loyaltyRepository.DeleteCustomer(user.UUID)
}

func TestCreateNotNewUser(t *testing.T) {

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&repositories.TwitchUser{})

	oldUser := &repositories.TwitchUser{
		UUID:       "existing-uuid",
		TwitchId:   "1234",
		TwitchNick: "My_Name",
	}
	db.Create(oldUser)

	settings := config.Config{
		LoyaltyServiceURI: "http://localhost:8081",
	}

	userService := NewPerfilService(&settings, db)

	user := twitch.User{
		ID:          "1234",
		DisplayName: "My_Name",
		Name:        "my_name",
	}

	msg, err := userService.CreateNewUser(user)
	assert.Equal(t, "My_Name usuário já existente, pare!", msg)
	assert.EqualError(t, err, errors.ErrUsuarioExistente.Error())

}

func TestGetUserCubes(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&repositories.TwitchUser{})

	settings := config.Config{
		LoyaltyServiceURI: "http://localhost:8081",
	}

	perfilService := NewPerfilService(&settings, db)

	twitchUser := twitch.User{
		ID:          "1234",
		DisplayName: "My_Name",
		Name:        "my_name",
	}
	msg, err := perfilService.CreateNewUser(twitchUser)
	assert.Equal(t, "My_Name usuário criado com sucesso", msg)
	assert.NoError(t, err)

	msg, err = perfilService.GetUserCubes(twitchUser)
	assert.Equal(t, "My_Name você tem 0 cubos", msg)
	assert.NoError(t, err)

	user, _ := perfilService.userRepository.GetUserByField("twitch_id", twitchUser.ID)
	perfilService.loyaltyRepository.DeleteCustomer(user.UUID)

}

func TestGetUserCubesNotExists(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&repositories.TwitchUser{})

	settings := config.Config{
		LoyaltyServiceURI: "http://localhost:8081",
	}

	userService := NewPerfilService(&settings, db)

	twitchUser := twitch.User{
		ID:          "1234",
		DisplayName: "My_Name",
		Name:        "my_name",
	}

	msg, err := userService.GetUserCubes(twitchUser)
	assert.Equal(t, "My_Name usuário não encontrado. Dê !join", msg)
	assert.EqualError(t, err, gorm.ErrRecordNotFound.Error())

}
