package services

import (
	"fmt"
	"log"
	"teomebot/config"
	"teomebot/errors"
	"teomebot/repositories"

	"github.com/gempir/go-twitch-irc/v4"
	"gorm.io/gorm"
)

type PerfilService struct {
	loyaltyRepository *repositories.LoyaltyRepository
	userRepository    *repositories.UserRepository
	retroRepository   *repositories.RetroRepository
}

func (s *PerfilService) CreateNewUser(twitchUser twitch.User) (string, error) {
	_, err := s.userRepository.GetUserByField("twitch_id", twitchUser.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			uuid, err := s.loyaltyRepository.CreateCustomerByTwitch(twitchUser.ID)
			if err != nil {
				log.Println(err)
				msg := fmt.Sprintf("%s não foi possível criar seu usuário", twitchUser.DisplayName)
				return msg, err
			}

			newTwitchUser := &repositories.TwitchUser{
				UUID:       uuid,
				TwitchId:   twitchUser.ID,
				TwitchNick: twitchUser.DisplayName,
			}

			if err := s.userRepository.CreateUser(newTwitchUser); err != nil {
				s.loyaltyRepository.DeleteCustomer(uuid)
				log.Println(err)
				msg := fmt.Sprintf("%s não foi possível criar seu usuário", twitchUser.DisplayName)
				return msg, err
			}

			msg := fmt.Sprintf("%s usuário criado com sucesso", twitchUser.DisplayName)
			return msg, nil
		}

		log.Println(err)
		msg := fmt.Sprintf("%s não foi possível criar seu usuário", twitchUser.DisplayName)
		return msg, err
	}

	msg := fmt.Sprintf("%s usuário já existente, pare!", twitchUser.DisplayName)
	return msg, errors.ErrUsuarioExistente
}

func (s *PerfilService) GetUserCubes(twitchUser twitch.User) (string, error) {

	user, err := s.userRepository.GetUserByField("twitch_id", twitchUser.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			msg := fmt.Sprintf("%s usuário não encontrado. Dê !join", twitchUser.DisplayName)
			return msg, err
		}
	}

	customer, err := s.loyaltyRepository.GetCustomer(user.UUID)
	if err != nil {
		log.Println(err)
		msg := fmt.Sprintf("%s não foi possível recuperar seus cubos", twitchUser.DisplayName)
		return msg, err
	}

	msg := fmt.Sprintf("%s você tem %d cubos", twitchUser.DisplayName, customer.NrPoints)
	return msg, nil

}

func (s *PerfilService) GetUserRetro(twitchUser twitch.User) (string, error) {

	user, err := s.userRepository.GetUserByField("twitch_id", twitchUser.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			msg := fmt.Sprintf("%s usuário não encontrado. Dê !join", twitchUser.DisplayName)
			return msg, err
		}
	}

	retro, err := s.retroRepository.GetUserRetro(user.UUID, user.TwitchNick)
	return *retro, err

}

func NewPerfilService(settings *config.Config, db *gorm.DB) *PerfilService {

	loyaltyRepository := repositories.NewLoyaltyRepository(settings)
	userRepository := repositories.NewUserRepository(db)

	return &PerfilService{
		loyaltyRepository: loyaltyRepository,
		userRepository:    userRepository,
	}

}
