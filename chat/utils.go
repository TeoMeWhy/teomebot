package chat

import (
	"log"
	"math/rand"
	"strings"
	"teomebot/controllers"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
)

// ParseCommand realiza o Parse do comando
func ParseCommand(m string) string {
	command := strings.TrimLeft(m, "!")
	command = strings.Split(command, " ")[0]
	command = strings.ToLower(command)
	return command
}

// GetChat fica recebendo os comandos do chat
func GetChat(client *twitch.Client, channel string) {

	log.Println("Ligando Bot!")

	allFun := map[string]HandleCommand{
		"agenda":     Agenda,
		"amazon":     Amazon,
		"anaconda":   Anaconda,
		"apoio":      Apoio,
		"asw":        ASW,
		"asn":        Asn,
		"ban":        Ban,
		"blog":       Blog,
		"caixa":      Caixa,
		"coach":      Coach,
		"comu":       Comunidade,
		"comuna":     Comunidade,
		"comunidade": Comunidade,
		"cubos":      Cubos,
		"cursos":     Cursos,
		"go":         GoCmd,
		"git":        Git,
		"github":     GitHub,
		"insta":      Instagram,
		"ig":         Instagram,
		"instagram":  Instagram,
		"join":       Join,
		"linkedin":   LinkedIn,
		"linuxtips":  LinuxTips,
		"loja":       Loja,
		"mega":       Mega,
		"metal":      Metal,
		"news":       News,
		"niver":      Niver,
		"pandas":     Pandas,
		"pdi":        PDI,
		"pix":        Pix,
		"ppt":        PPT,
		"presente":   Presente,
		"prime":      Prime,
		"projeto":    Projeto,
		"fidelidade": Fidelidade,
		"refs":       Refs,
		"retro":      Retro,
		"rock":       Rock,
		"sql":        SQL,
		"sub":        Sub,
		"troca":      Troca,
		"twitter":    Twitter,
		"vscode":     VSCode,
		"x":          Twitter,
		"youtube":    YouTube,
		"nekt":       Nekt,
	}

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {

		if strings.HasPrefix(message.Message, "!") {
			command := ParseCommand(message.Message)
			if execCommand, ok := allFun[command]; ok {
				go execCommand(client, message)
			}
		} else {
			go controllers.MessageChatController(message.User)
		}
	})

	client.Join(channel)
	err := client.Connect()
	if err != nil {
		panic(err)
	}
}

// RandomWarnings envia avisos no chat
func RandomWarnings(client *twitch.Client, channel string) {

	log.Println("Ligando avisos!")

	warnings := []HandleCommand{
		Apoio,
		ASW,
		Caixa,
		Comunidade,
		News,
		Pix,
		Prime,
		Sub,
		Sub,
		Niver,
	}

	client.Join(channel)
	rand.Seed(time.Now().UnixNano())
	for {
		choice := warnings[rand.Intn(len(warnings))]
		choice(client, twitch.PrivateMessage{Channel: channel})
		time.Sleep(time.Minute * time.Duration(rand.Float32()*10+5))
	}
}
