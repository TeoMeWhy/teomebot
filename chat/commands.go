package chat

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"teomebot/controllers"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
)

// NetLink envia o link da net com parâmetros de campanha
func Nekt(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Conheça a Nekt: https://nekt.com/?via=33hoqj8m"
	c.Say(m.Channel, text)
}

// HandleCommand é um tipo para ser passado em um mapa de valores
type HandleCommand func(*twitch.Client, twitch.PrivateMessage)

// Agenda Mostra nossa agenda de cursos futuros
func Agenda(c *twitch.Client, m twitch.PrivateMessage) {
	text := `Confira nossa agenda de próximos cursos: https://cursos.teomewhy.org/material_2025`
	c.Say(m.Channel, text)
}

func Amazon(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Aproveite a Black Friday da Amazon: https://amzn.to/49Hh20u"
	c.Say(m.Channel, text)
}

// Asn exibe no chat o link de parceira com Adriana Silva
func Apoio(c *twitch.Client, m twitch.PrivateMessage) {
	text := `Financie nosso projeto:   Pix.....................pix@teomewhy.org ApoiaSe.............apoia.se/teomewhy LivePix.............livepix.gg/teomewhy`
	c.Say(m.Channel, text)
}

// Asn exibe no chat o link de parceira com Adriana Silva
func Asn(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Cursos maravilhosos com a Jedi em Analytics: https://asn.rocks/"
	c.Say(m.Channel, text)
}

// ASW exibe o site do Instituto Aaron Swartz
func ASW(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Conheça o Instituto Aaron Swartz: institutoasw.org"
	c.Say(m.Channel, text)
}

// Ban manda mensagem de manimento
func Ban(c *twitch.Client, m twitch.PrivateMessage) {

	text := ""
	sliceBan := strings.Split(m.Message, " ")
	if len(sliceBan) == 1 {
		text = fmt.Sprintf("%s foi banido por não saber usar o comando de ban!", m.User.Name)
	} else {
		userBan := sliceBan[1]
		frases := []string{
			"não saber cantar Téo Me Why.",
			"não ter sorte no dado.",
			"usar base de treino como validação.",
			"fazer não fazer análise exploratória.",
			"não fazer análise de resíduo.",
			"confiar apenas acurácia.",
			"fazer UPDATE sem WHERE.",
			"dar DROP DATABASE.",
			"fazer DELETE sem WHERE.",
			"fazer amostra viciada.",
			"não saber fazer JOIN.",
			"colocar Jupyter Notebook em produção.",
			"entregar o modelo em PPT.",
			"ser um Engenheiro de Dados Quântico.",
		}
		rand.Seed(time.Now().UnixNano())
		choice := frases[rand.Intn(len(frases))]
		text = fmt.Sprintf("%s baniu %s por %s", m.User.Name, userBan, choice)
	}
	c.Say(m.Channel, text)
}

// Blog fala do nosso blog
func Blog(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Conheça nosso blog: https://teomewhy.org"
	c.Say(m.Channel, text)
}

func Caixa(c *twitch.Client, m twitch.PrivateMessage) {
	txt := "Destinatário: Téo Calvo - R. João Batista Colnago, 151 - Vila Liberdade, Pres. Prudente-SP CEP 19050-970 CAIXA POSTAL 3094"
	c.Say(m.Channel, txt)
}

// Coach retorna uma frase motivacional para o chat
func Coach(c *twitch.Client, m twitch.PrivateMessage) {
	frases := []string{
		"%s, no tempo certo, tudo dará errado.",
		"%s, acorda, o fracasso te espera.",
		"%s, nunca subestime sua incapacidade.",
		"%s, seja forte, desista!",
		"%s, fracasse enquanto eles descansam.",
		"%s, a vida te derruba hj, preparando para a queda de amanhã.",
		"%s, nunca é tarde para desistir.",
		"%s, você é mais fraco do que pensa.",
		"%s, nada vem no tempo certo.",
		"%s, todo fracasso começa com a decisão de tentar.",
		"%s, você só perderá amanhã se não desistir hoje.",
		"%s, desistir é sempre a melhor opção.",
		"%s, descubra novas formas de fracassar.",
		"%s, a expectativa é mãe da merda!",
		"%s, nunca aceite críticas construitivas de quem nunca construiu nada.",
		"%s, continue sonhando, você acabará se atrasando",
	}

	rand.Seed(time.Now().UnixNano())
	choice := fmt.Sprintf(frases[rand.Intn(len(frases))], m.User.Name)
	c.Say(m.Channel, choice)
}

func Comandos(c *twitch.Client, m twitch.PrivateMessage) {
	text := `!agenda !apoio !asw !asn !blog !cursos !git !github !linuxtips !loja !metal !pdi !pix !pontos !refs !rock !sql !cubos !join !presente !troca`
	c.Say(m.Channel, text)
}

func Comunidade(c *twitch.Client, m twitch.PrivateMessage) {
	text := `Entre para nossa comunidade: comunidade.teomewhy.org`
	c.Say(m.Channel, text)
}

func Cursos(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Plataforma de cursos livres: https://cursos.teomewhy.org"
	c.Say(m.Channel, text)
}

func Fidelidade(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Entenda nosso sistema de pontos: https://teomewhy.org/twitch#sistema-de-pontos"
	c.Say(m.Channel, text)
}

// GitHub exibe nosso perfil no GitHub
func GitHub(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Se liga no meu GitHub: https://github.com/TeoMeWhy"
	c.Say(m.Channel, text)
}

// Instagram exibe nossp p[erfil do instagram
func Instagram(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Me siga no Instagram: https://www.instagram.com/teomewhy"
	c.Say(m.Channel, text)
}

// LinkedIn exibe nosso perfil do linkedin
func LinkedIn(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Me adicione no LinkedIn https://www.linkedin.com/in/teocalvo/"
	c.Say(m.Channel, text)
}

// LinkedIn exibe nosso perfil do linkedin
func LinuxTips(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Conheça a LinuxTips e seus cursos fodas! https://www.linuxtips.io/home"
	c.Say(m.Channel, text)
}

func Loja(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Acesse nossa lojinha para resgate de prêmios: https://streamelements.com/teomewhy/store"
	c.Say(m.Channel, text)
}

// Metal mostra link da amazonPrime e Twitch
func Metal(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Playlist Metal: https://open.spotify.com/playlist/2EyQ31SxCDDEdn2sRrMGQX?si=2qJpNrnTRW6dyhmiSLiiTg"
	c.Say(m.Channel, text)
}

// News mostra link da nossa newsletter
func News(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Se inscreva na nossa newsletter: https://teomewhy.substack.com"
	c.Say(m.Channel, text)
}

// PDI mostra o link da planilha de PDI
func PDI(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Confira nosso vídeo de PDI: https://youtu.be/L0G07W5aODM"
	c.Say(m.Channel, text)
}

// Pix mostra o link para nosso livepix
func Pix(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Nos envie uma mensagem pelo livePix: https://livepix.gg/teomewhy"
	c.Say(m.Channel, text)
}

func Projeto(c *twitch.Client, m twitch.PrivateMessage) {
	txt := "Estamos trabalhando com os dados da Formula 1."
	c.Say(m.Channel, txt)
}

// Prime mostra link da amazonPrime e Twitch
func Prime(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Vincule seu Amazon Prime com a Twitch e apoie nosso projeto!! https://twitch.amazon.com/tp"
	c.Say(m.Channel, text)
}

// Refs mostra a lista de referencia que temos em DataScience
func Refs(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Lista completas de referências em Data Science, Programação e Estatística: https://github.com/TeoMeWhy/teomerefs"
	c.Say(m.Channel, text)
}

// Rock Playlist de Rock'n Roll
func Rock(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Playlist Rock: https://open.spotify.com/playlist/4PMBaBW2WAXexrBjbd9LlX?si=efb5a9d8346e4de1"
	c.Say(m.Channel, text)
}

func Twitter(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Me siga no Twitter ou X:  https://x.com/teomewhy"
	c.Say(m.Channel, text)
}

// YouTube exibe nosso canal no YouTube
func YouTube(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Se inscreva no nosso canal do YouTube: https://www.youtube.com/@teomewhy"
	c.Say(m.Channel, text)
}

// -------- //

func Anaconda(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Faça o download da Anaconda aqui: https://www.anaconda.com/download"
	c.Say(m.Channel, text)
}

func VSCode(c *twitch.Client, m twitch.PrivateMessage) {
	text := "Faça o download do Visual Studio Code aqui: https://code.visualstudio.com/download"
	c.Say(m.Channel, text)
}

// ----------- //

func Mega(c *twitch.Client, m twitch.PrivateMessage) {
	msg := controllers.MegaController(m.User)
	c.Say(m.Channel, msg)
}

// ----------- //

func Join(c *twitch.Client, m twitch.PrivateMessage) {

	err := controllers.ExecCreateOrUpdateUser(&m.User)
	if err != nil {

		if err.Error() == "usuário sem atualização" {
			txt := "%s você acaba de perder 100 pontos por usar !join indevidamente"
			c.Say(m.Channel, fmt.Sprintf(txt, m.User.Name))
			return
		}

		log.Println(err)
		c.Say(m.Channel, fmt.Sprintf("%s não foi possível criar seu usuário", m.User.Name))
		return
	}

	c.Say(m.Channel, fmt.Sprintf("%s usuário criado ou atualizado com sucesso", m.User.Name))
}

func Cubos(c *twitch.Client, m twitch.PrivateMessage) {
	msg := controllers.GetUserPointsController(m.User)
	c.Say(m.Channel, msg)
}

func Presente(c *twitch.Client, m twitch.PrivateMessage) {
	msg := controllers.PresentController(m.User)
	if msg == "" {
		return
	}

	c.Say(m.Channel, msg)
}

func Profile(c *twitch.Client, m twitch.PrivateMessage) {
	msg := controllers.ProfileController(m.User)
	c.Say(m.Channel, msg)
}

func Troca(c *twitch.Client, m twitch.PrivateMessage) {
	msg := controllers.TrocaController(m.User)
	c.Say(m.Channel, msg)
}

func Retro(c *twitch.Client, m twitch.PrivateMessage) {
	msg := controllers.RetroController(m.User)
	c.Say(m.Channel, msg)
}
