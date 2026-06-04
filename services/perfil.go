package services

import (
	"fmt"
	"log"
	"net/http"
	"teomebot/config"
	"teomebot/errors"
	"teomebot/models"
	"teomebot/repositories"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
	"gorm.io/gorm"
)

type PerfilService struct {
	loyaltyRepository  *repositories.LoyaltyRepository
	userRepository     *repositories.UserRepository
	palantirRepository *repositories.PalantirRepository
}

func (s *PerfilService) CreateNewUser(twitchUser twitch.User) (string, error) {

	_, err := s.loyaltyRepository.GetCustomerByTwitch(twitchUser.ID)
	if err != nil {

		_, err = s.loyaltyRepository.CreateCustomerByTwitch(twitchUser.ID)
		if err != nil {
			log.Println(err)
			msg := fmt.Sprintf("%s não foi possível criar seu usuário", twitchUser.DisplayName)
			return msg, err
		}

		msg := fmt.Sprintf("%s usuário criado com sucesso. Aproveite para conhecer nossas trilhas de conhecimento: cursos.teomewhy.org", twitchUser.DisplayName)
		return msg, nil

	}

	msg := fmt.Sprintf("%s usuário já existente, pare!", twitchUser.DisplayName)
	return msg, errors.ErrUsuarioExistente
}

func (s *PerfilService) GetUserCubes(twitchUser twitch.User) (string, error) {

	customer, err := s.loyaltyRepository.GetCustomerByTwitch(twitchUser.ID)
	if err != nil {
		log.Println(err)
		msg := fmt.Sprintf("%s não foi possível recuperar seus cubos", twitchUser.DisplayName)
		return msg, err
	}

	if customer.DescCustomerName != twitchUser.DisplayName {
		customer.DescCustomerName = twitchUser.DisplayName
		if err := s.loyaltyRepository.UpdateCustomer(*customer); err != nil {
			log.Println(err)
		}
	}

	msg := fmt.Sprintf("%s você tem %d cubos", twitchUser.DisplayName, customer.NrPoints)
	return msg, nil

}

func (s *PerfilService) GetFielScore(twitchUser twitch.User) (string, error) {

	user, err := s.loyaltyRepository.GetCustomerByTwitch(twitchUser.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			msg := fmt.Sprintf("%s usuário não encontrado. Dê !join", twitchUser.DisplayName)
			return msg, err
		}
		msg := fmt.Sprintf("%s não foi possível obter seu fiel score", twitchUser.DisplayName)
		return msg, err
	}

	predict, err := s.palantirRepository.GetPrediction("tmw_score_fiel", user.UUID)
	if err != nil {
		msg := fmt.Sprintf("%s não foi possível obter seu fiel score", twitchUser.DisplayName)
		return msg, err
	}

	score, ok := predict["score_fiel"]
	if !ok {
		msg := fmt.Sprintf("%s não foi possível obter seu fiel score", twitchUser.DisplayName)
		return msg, errors.ErrFielScoreNotFound
	}

	score *= 100

	checkFiel, err := s.CheckFielToday(twitchUser)
	if err != nil {
		log.Println(err.Error())
	}

	var msg string
	if score < 10 {
		msg = fmt.Sprintf("%s seu Fiel-Score: %.2f%%. Isso é um tanto quanto vergonhoso! Bora interagir mais.", twitchUser.DisplayName, score)
	} else if score < 25 {
		msg = fmt.Sprintf("%s seu Fiel-Score: %.2f%%. Dá para melhorar bastante!", twitchUser.DisplayName, score)
	} else if score < 50 {
		msg = fmt.Sprintf("%s seu Fiel-Score: %.2f%%. Está no caminho certo para começar a ganhar recompensas!", twitchUser.DisplayName, score)
	} else if score < 75 {
		msg = fmt.Sprintf("%s seu Fiel-Score: %.2f%%. Muito bom! Continue assim para ganhar recompensas cada vez melhores!", twitchUser.DisplayName, score)
	} else {
		msg = fmt.Sprintf("%s seu Fiel-Score: %.2f%%. Parabéns, você é um dos fiéis da comunidade! Exemplo a ser seguido!", twitchUser.DisplayName, score)
	}

	if checkFiel {
		return msg, nil
	}

	prodFiel := models.NewFielScore(score)
	products := []models.ProductPoints{prodFiel}
	if err := s.loyaltyRepository.AddPoints(user.UUID, products); err != nil {
		log.Println(err)
		msg := fmt.Sprintf("%s não foi possível adicionar seus pontos de fiel", twitchUser.DisplayName)
		return msg, err
	}

	var msgPontos string
	if prodFiel.VlProduct > 0 {
		msgPontos = fmt.Sprintf(" Você ganhou %d cubos!", prodFiel.VlProduct)
	} else {
		msgPontos = fmt.Sprintf(" Você perdeu %d cubos!", -prodFiel.VlProduct)
	}

	msg += msgPontos

	customer, err := s.loyaltyRepository.GetCustomerByTwitch(twitchUser.ID)
	if err != nil {
		log.Println(err)
	}

	if customer.DescCustomerName != twitchUser.DisplayName {
		customer.DescCustomerName = twitchUser.DisplayName
		if err := s.loyaltyRepository.UpdateCustomer(*customer); err != nil {
			log.Println(err)
		}
	}

	return msg, nil
}

func (s *PerfilService) CheckFielToday(twitchUser twitch.User) (bool, error) {

	user, err := s.loyaltyRepository.GetCustomerByTwitch(twitchUser.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, err
		}
		return false, err
	}

	lastDate, err := s.loyaltyRepository.GetCustomerLastTransactionDateByCategory(user.UUID, "fiel")
	if err != nil {
		return false, err
	}

	if lastDate == nil {
		return false, nil
	}

	today := time.Now().Truncate(24 * time.Hour)
	lastTransactionDate := lastDate.Truncate(24 * time.Hour)

	return today.Equal(lastTransactionDate), nil
}

func NewPerfilService(settings *config.Config, db *gorm.DB) *PerfilService {

	loyaltyRepository := repositories.NewLoyaltyRepository(settings)
	palantirRepository := repositories.NewPalantirRepository(settings, &http.Client{})
	userRepository := repositories.NewUserRepository(db)

	return &PerfilService{
		loyaltyRepository:  loyaltyRepository,
		userRepository:     userRepository,
		palantirRepository: palantirRepository,
	}

}
