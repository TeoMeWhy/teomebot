package services

import (
	"fmt"
	"log"
	"strings"
	"teomebot/config"
	"teomebot/repositories"

	"github.com/gempir/go-twitch-irc/v4"
)

type IAService struct {
	ragiaClient *repositories.RagiaClient
}

func NewIAService(cfg *config.Config) *IAService {

	ragiaClient := repositories.NewRagiaClient(cfg)

	return &IAService{
		ragiaClient: ragiaClient,
	}
}

func (s *IAService) GetAIResponse(msg twitch.PrivateMessage) (string, error) {

	log.Println("Validando mensagem para a AI")

	if !strings.ContainsAny(msg.Message, "?") {
		return "", nil
	}

	log.Println("Mensagem válida para enviar para AI")

	response, err := s.ragiaClient.GetQueryResponse(msg.Message)
	if err != nil || response == "" {
		log.Println("Erro ou tipo de pergunta fora de contexto")
		return "", err
	}

	txt := fmt.Sprintf("%s %s", msg.User.DisplayName, response)
	log.Println(txt)
	return txt, nil
}
