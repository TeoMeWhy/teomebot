package controllers

import (
	"teomebot/repositories"

	"gorm.io/gorm"
)

var mensagens = []repositories.Messagem{

	{
		Chave:    "agenda",
		Conteudo: "Confira nossa agenda de próximos cursos: https://cursos.teomewhy.org/material_2025",
	},

	{
		Chave:    "apoio",
		Conteudo: "Financie nosso projeto:   Pix.....................pix@teomewhy.org ApoiaSe.............apoia.se/teomewhy LivePix.............livepix.gg/teomewhy"},

	{
		Chave:    "comunidade",
		Conteudo: "Entre para nossa comunidade: comunidade.teomewhy.org"},

	{
		Chave:    "nekt",
		Conteudo: "Conheça a Nekt: https://nekt.com/?via=33hoqj8m",
	},

	{
		Chave:    "amazon",
		Conteudo: "Aproveite a Black Friday da Amazon: https://amzn.to/49Hh20u",
	},

	{
		Chave:    "asn",
		Conteudo: "Cursos maravilhosos com a Jedi em Analytics: https://asn.rocks/",
	},

	{
		Chave:    "asw",
		Conteudo: "Conheça o Instituto Aaron Swartz: institutoasw.org",
	},
	{
		Chave:    "blog",
		Conteudo: "Conheça nosso blog: https://teomewhy.org",
	},

	{
		Chave:    "caixa",
		Conteudo: "Destinatário: Téo Calvo - R. João Batista Colnago, 151 - Vila Liberdade, Pres. Prudente-SP CEP 19050-970 CAIXA POSTAL 3094",
	},

	{
		Chave:    "cursos",
		Conteudo: "Plataforma de cursos livres: https://cursos.teomewhy.org",
	},

	{
		Chave:    "fidelidade",
		Conteudo: "Entenda nosso sistema de pontos: https://teomewhy.org/twitch#sistema-de-pontos",
	},

	{
		Chave:    "github",
		Conteudo: "Se liga no meu GitHub: https://github.com/TeoMeWhy",
	},

	{
		Chave:    "instagram",
		Conteudo: "Me siga no Instagram: https://www.instagram.com/teomewhy",
	},

	{
		Chave:    "linkedin",
		Conteudo: "Me adicione no LinkedIn https://www.linkedin.com/in/teocalvo/",
	},

	{
		Chave:    "linuxtips",
		Conteudo: "Conheça a LinuxTips e seus cursos fodas! https://www.linuxtips.io/home",
	},

	{
		Chave:    "loja",
		Conteudo: "Loja na Shopee: https://shopee.com.br/ferhsilvestre Resgate seu cupom aqui: https://streamelements.com/teomewhy/store",
	},

	{
		Chave:    "metal",
		Conteudo: "Playlist Metal: https://open.spotify.com/playlist/2EyQ31SxCDDEdn2sRrMGQX?si=2qJpNrnTRW6dyhmiSLiiTg",
	},

	{
		Chave:    "news",
		Conteudo: "Se inscreva na nossa newsletter: https://teomewhy.substack.com",
	},

	{
		Chave:    "pdi",
		Conteudo: "Confira nosso vídeo de PDI: https://youtu.be/L0G07W5aODM",
	},

	{
		Chave:    "pix",
		Conteudo: "Nos envie uma mensagem pelo livePix: https://livepix.gg/teomewhy",
	},

	{
		Chave:    "projeto",
		Conteudo: "Estamos trabalhando com os dados da Formula 1.",
	},

	{
		Chave:    "prime",
		Conteudo: "Vincule seu Amazon Prime com a Twitch e apoie nosso projeto!! https://twitch.amazon.com/tp",
	},

	{
		Chave:    "refs",
		Conteudo: "Lista completas de referências em Data Science, Programação e Estatística: https://github.com/TeoMeWhy/teomerefs",
	},

	{
		Chave:    "rock",
		Conteudo: "Playlist Rock: https://open.spotify.com/playlist/4PMBaBW2WAXexrBjbd9LlX?si=efb5a9d8346e4de1",
	},

	{
		Chave:    "twitter",
		Conteudo: "Me siga no Twitter ou X:  https://x.com/teomewhy",
	},

	{
		Chave:    "youtube",
		Conteudo: "Se inscreva no nosso canal do YouTube: https://www.youtube.com/@teomewhy",
	},

	{
		Chave:    "anaconda",
		Conteudo: "Faça o download da Anaconda aqui: https://www.anaconda.com/download",
	},

	{
		Chave:    "vscode",
		Conteudo: "Faça o download do Visual Studio Code aqui: https://code.visualstudio.com/download",
	},

	{
		Chave:    "workshop",
		Conteudo: "Não perca o workshop da Nekt! https://materiais.nekt.com/workshop-pyspark-02-26?utm_medium=influencer&utm_source=twitch&utm_campaign=tmw&utm_content=live_pyspark",
	},
}

func SetMensagens(db *gorm.DB) {
	db.AutoMigrate(&repositories.Messagem{})
	db.Save(&mensagens)
}
