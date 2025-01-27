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
		"apoio":      Apoio,
		"asw":        ASW,
		"asn":        Asn,
		"ban":        Ban,
		"blog":       Blog,
		"caixa":      Caixa,
		"coach":      Coach,
		"cursos":     Cursos,
		"go":         GoCmd,
		"git":        Git,
		"github":     GitHub,
		"insta":      Instagram,
		"ig":         Instagram,
		"instagram":  Instagram,
		"linkedin":   LinkedIn,
		"linuxtips":  LinuxTips,
		"loja":       Loja,
		"metal":      Metal,
		"news":       News,
		"niver":      Lista,
		"pdi":        PDI,
		"pix":        Pix,
		"prime":      Prime,
		"projeto":    Projeto,
		"fidelidade": Fidelidade,
		"refs":       Refs,
		"rock":       Rock,
		"sql":        SQL,
		"sub":        Sub,

		"telegram": Telegram,
		"twitter":  Twitter,
		"youtube":  YouTube,

		"pandas":   Pandas,
		"anaconda": Anaconda,
		"vscode":   VSCode,

		"mega": Mega,

		"join":     Join,
		"cubos":    Cubos,
		"presente": Presente,
		// "profile":  Profile,
		"troca": Troca,
		"retro": Retro,

		"ppt": PPT,
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
		News,
		Pix,
		Prime,
		Sub,
		Sub,
	}

	client.Join(channel)
	rand.Seed(time.Now().UnixNano())
	for {
		choice := warnings[rand.Intn(len(warnings))]
		choice(client, twitch.PrivateMessage{Channel: channel})
		time.Sleep(time.Minute * time.Duration(rand.Float32()*10+5))
	}
}
