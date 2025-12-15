package controllers

import (
	"teomebot/config"
	"testing"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTest() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	SetMensagens(db)
	return db

}

func TestGetMessageByCommand(t *testing.T) {

	db := setupTest()
	settings := &config.Config{
		LoyaltyServiceURI: "http://localhost:8081",
		RetroServiceURI:   "http://localhost:8082",
	}

	controller, err := NewCommandsController(nil, db, settings)
	assert.NoError(t, err)

	testsList := []struct {
		command       string
		expectedReply string
	}{
		{
			"agenda",
			"Confira nossa agenda de próximos cursos: https://cursos.teomewhy.org/material_2025",
		},
		{
			"apoio",
			"Financie nosso projeto:   Pix.....................pix@teomewhy.org ApoiaSe.............apoia.se/teomewhy LivePix.............livepix.gg/teomewhy",
		},
		{
			"comunidade",
			"Entre para nossa comunidade: comunidade.teomewhy.org",
		},
		{
			"nekt",
			"Conheça a Nekt: https://nekt.com/?via=33hoqj8m",
		},
		{
			"amazon",
			"Aproveite a Black Friday da Amazon: https://amzn.to/49Hh20u",
		},
		{
			"asn",
			"Cursos maravilhosos com a Jedi em Analytics: https://asn.rocks/",
		},
		{
			"asw",
			"Conheça o Instituto Aaron Swartz: institutoasw.org",
		},
		{
			"blog",
			"Conheça nosso blog: https://teomewhy.org",
		},
		{
			"caixa",
			"Destinatário: Téo Calvo - R. João Batista Colnago, 151 - Vila Liberdade, Pres. Prudente-SP CEP 19050-970 CAIXA POSTAL 3094",
		},
		{
			"cursos",
			"Plataforma de cursos livres: https://cursos.teomewhy.org",
		},
		{
			"fidelidade",
			"Entenda nosso sistema de pontos: https://teomewhy.org/twitch#sistema-de-pontos",
		},
		{
			"github",
			"Se liga no meu GitHub: https://github.com/TeoMeWhy",
		},
		{
			"instagram",
			"Me siga no Instagram: https://www.instagram.com/teomewhy",
		},
		{
			"linkedin",
			"Me adicione no LinkedIn https://www.linkedin.com/in/teocalvo/",
		},
		{
			"linuxtips",
			"Conheça a LinuxTips e seus cursos fodas! https://www.linuxtips.io/home",
		},
		{
			"loja",
			"Acesse nossa lojinha para resgate de prêmios: https://streamelements.com/teomewhy/store",
		},
		{
			"metal",
			"Playlist Metal: https://open.spotify.com/playlist/2EyQ31SxCDDEdn2sRrMGQX?si=2qJpNrnTRW6dyhmiSLiiTg",
		},
		{
			"news",
			"Se inscreva na nossa newsletter: https://teomewhy.substack.com",
		},
		{
			"pdi",
			"Confira nosso vídeo de PDI: https://youtu.be/L0G07W5aODM",
		},
		{
			"pix",
			"Nos envie uma mensagem pelo livePix: https://livepix.gg/teomewhy",
		},
		{
			"projeto",
			"Estamos trabalhando com os dados da Formula 1.",
		},
		{
			"prime",
			"Vincule seu Amazon Prime com a Twitch e apoie nosso projeto!! https://twitch.amazon.com/tp",
		},
		{
			"refs",
			"Lista completas de referências em Data Science, Programação e Estatística: https://github.com/TeoMeWhy/teomerefs",
		},
		{
			"rock",
			"Playlist Rock: https://open.spotify.com/playlist/4PMBaBW2WAXexrBjbd9LlX?si=efb5a9d8346e4de1",
		},
		{
			"twitter",
			"Me siga no Twitter ou X:  https://x.com/teomewhy",
		},
		{
			"youtube",
			"Se inscreva no nosso canal do YouTube: https://www.youtube.com/@teomewhy",
		},
		{
			"anaconda",
			"Faça o download da Anaconda aqui: https://www.anaconda.com/download",
		},
		{
			"vscode",
			"Faça o download do Visual Studio Code aqui: https://code.visualstudio.com/download",
		},
		{
			"NoCommand",
			"",
		},
	}

	for _, aTest := range testsList {
		t.Run(aTest.command, func(t *testing.T) {

			twitchMessage := twitch.PrivateMessage{
				Message: "!" + aTest.command,
				User:    twitch.User{DisplayName: "test_user"},
			}

			msg, err := controller.GetCommand(twitchMessage)
			assert.Equal(t, aTest.expectedReply, msg)
			assert.NoError(t, err)
		})
	}

}
