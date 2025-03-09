package main

import (
	"os"
	"teomebot/chat"
	"teomebot/models"
	"teomebot/utils"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	user := os.Getenv("TWITCH_BOT")
	oauth := os.Getenv("TWITCH_OAUTH_BOT")
	channel := os.Getenv("TWITCH_CHANNEL")

	con, err := utils.OpenDBConnection()
	if err != nil {
		panic("erro na conexão com banco")
	}

	con.AutoMigrate(&models.PresentUser{}, &models.StreakPresentUser{}, &models.TwitchUser{}, &models.ProfileUser{})

	client := twitch.NewClient(user, oauth)

	go chat.GetChat(client, channel)
	go chat.RandomWarnings(client, channel)

	for {
		time.Sleep(time.Hour * 1)
	}
}
