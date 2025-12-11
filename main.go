package main

import (
	"log"
	"teomebot/config"
	"teomebot/controllers"
	"teomebot/utils"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	settings, err := config.LoadConfig()
	if err != nil {
		log.Println(err)
		panic("erro ao carregar configuração")
	}

	con, err := utils.OpenDBConnection(settings)
	if err != nil {
		panic("erro na conexão com banco")
	}

	log.Println("Criando cliente da Twitch")
	client := twitch.NewClient(settings.TwitchBot, settings.TwitchOauthBot)

	log.Println("Iniciando o controller")
	controllerMessages, err := controllers.NewCommandsController(client, con, settings)
	if err != nil {
		panic("erro ao iniciar controller de comandos")
	}

	log.Println("Capturando comandos")
	go controllerMessages.HandleCommands()

	for {
		time.Sleep(time.Hour * 1)
	}
}
