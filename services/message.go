package services

import (
	"teomebot/repositories"

	"gorm.io/gorm"
)

type MessageService struct {
	mensagensEstaticas  *repositories.MessageRepository
	mensagensAleatorias *repositories.MensagensAleatoriasRepository
}

func (s *MessageService) GetMensagem(key string) string {
	msg := s.mensagensEstaticas.ShowMessage(key)
	return msg
}

func NewMessageService(db *gorm.DB) *MessageService {

	mensagensEstaticas := repositories.NewMessageRepository(db)
	mensagensAleatorias := repositories.NewMensagensAleatoriasRepository()

	return &MessageService{
		mensagensEstaticas:  mensagensEstaticas,
		mensagensAleatorias: mensagensAleatorias,
	}

}
