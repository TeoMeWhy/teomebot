package services

import (
	"fmt"
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

	if !strings.ContainsAny(msg.Message, "?") {
		return "", nil
	}

	response, err := s.ragiaClient.GetQueryResponse(msg.Message)
	if err != nil || response == "" {
		return "", err
	}

	txt := fmt.Sprintf("%s %s", msg.User.DisplayName, response)
	return txt, nil
}
