package controllers

import (
	"log"
	"strings"
	"teomebot/config"
	"teomebot/services"

	"github.com/gempir/go-twitch-irc/v4"
	"gorm.io/gorm"
)

type CommandHandler func(twitch.Client, twitch.PrivateMessage)

type CommandsController struct {
	twitchChannel  string
	twitchClient   *twitch.Client
	perfilService  services.PerfilService
	pointsService  services.PointsService
	messageService services.MessageService
}

func (c *CommandsController) HandleMessages() {

	log.Println("Estou aqui!")

	c.twitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {

		if strings.HasPrefix(message.Message, "!") {
			msg, err := c.GetCommand(message)
			if err != nil {
				log.Println(err)
			}
			if msg != "" {
				c.twitchClient.Say(c.twitchChannel, msg)
			}

		} else {
			c.pointsService.AddMsgCubes(message.User)
		}

	})

	c.twitchClient.Join(c.twitchChannel)
	err := c.twitchClient.Connect()
	if err != nil {
		panic(err)
	}

}

func (c *CommandsController) GetCommand(message twitch.PrivateMessage) (string, error) {

	command := strings.Split(message.Message, " ")[0]
	command = strings.TrimPrefix(command, "!")

	switch command {

	case "join":
		return c.perfilService.CreateNewUser(message.User)

	case "cubos":
		return c.perfilService.GetUserCubes(message.User)

	case "retro":
		return c.perfilService.GetUserRetro(message.User)

	case "presente":
		return c.pointsService.MgmtPresenca(message.User)

	case "troca":
		return c.pointsService.CubesToDatapoints(message.User)

	default:
		msg := c.messageService.GetMensagem(command)
		return msg, nil
	}

}

func NewCommandsController(twitchclient *twitch.Client, db *gorm.DB, settings *config.Config) (*CommandsController, error) {

	perfilService := services.NewPerfilService(settings, db)
	pointsService := services.NewPointsService(settings, db)
	messageService := services.NewMessageService(db)

	controller := &CommandsController{
		twitchChannel:  settings.TwitchChannel,
		twitchClient:   twitchclient,
		perfilService:  *perfilService,
		pointsService:  *pointsService,
		messageService: *messageService,
	}

	SetMensagens(db)

	return controller, nil
}
