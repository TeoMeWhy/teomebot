package services

import (
	"fmt"
	"log"
	"teomebot/config"
	"teomebot/models"
	"teomebot/repositories"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
	"gorm.io/gorm"
)

type PointsService struct {
	loyaltyRepository        *repositories.LoyaltyRepository
	userRepository           *repositories.UserRepository
	presentRepository        *repositories.PresencaRepository
	streamElementsRepository *repositories.StreamElementsRepository
}

func NewPointsService(settings *config.Config, db *gorm.DB) *PointsService {
	loyaltyRepository := repositories.NewLoyaltyRepository(settings)
	userRepository := repositories.NewUserRepository(db)
	presentRepository := repositories.NewPresencaRepository(db)
	streamElementsRepository := repositories.NewStreamElementsRepository(settings)

	return &PointsService{
		loyaltyRepository:        loyaltyRepository,
		userRepository:           userRepository,
		presentRepository:        presentRepository,
		streamElementsRepository: streamElementsRepository,
	}

}

func (s *PointsService) AddMsgCubes(twitchUser twitch.User) {

	customer, err := s.loyaltyRepository.GetCustomerByTwitch(twitchUser.ID)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Println(err)
		}
		return
	}

	product := models.NewChatMessage()
	products := []models.ProductPoints{product}

	if err := s.loyaltyRepository.AddPoints(customer.UUID, products); err != nil {
		log.Println(err)
		return
	}

	if customer.DescCustomerName != twitchUser.DisplayName {
		customer.DescCustomerName = twitchUser.DisplayName
		if err := s.loyaltyRepository.UpdateCustomer(*customer); err != nil {
			log.Println(err, " - ", customer.UUID, " - ", twitchUser.DisplayName)
		}
	}

}

func (s *PointsService) addPresencaCubes(customer *repositories.Customer) error {

	products := []models.ProductPoints{models.NewPresent()}
	if err := s.loyaltyRepository.AddPoints(customer.UUID, products); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *PointsService) addStreakCubes(twitchID string) {

	streak, err := s.presentRepository.LoadLastUpdatedStreak(twitchID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := s.presentRepository.CreateStreak(twitchID); err != nil {
				log.Println(err)
				return
			}
			return
		}
		return
	}

	pastUpdate := streak.UpdatedAt.Year() + streak.UpdatedAt.YearDay()
	currentDate := time.Now().Year() + time.Now().YearDay()

	if currentDate-pastUpdate <= 1 {
		streak.Qtd += 1
		if err := s.presentRepository.UpdateStreak(streak); err != nil {
			log.Println(err)
			return
		}
		if streak.Qtd%5 == 0 {
			products := []models.ProductPoints{}
			for i := 0; i < int(streak.Qtd/5); i++ {
				products = append(products, models.NewStreakPresent())
			}
			if err := s.loyaltyRepository.AddPoints(twitchID, products); err != nil {
				log.Println(err)
				return
			}
		}
		return
	}

	if err := s.presentRepository.CreateStreak(twitchID); err != nil {
		log.Println(err)
		return
	}

}

func (s *PointsService) CheckPresentToday(id string) (bool, error) {

	lastDate, err := s.loyaltyRepository.GetCustomerLastTransactionDateByCodProduct(id, "11")
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

func (s *PointsService) MgmtPresenca(twitchUser twitch.User) (string, error) {

	customer, err := s.loyaltyRepository.GetCustomerByTwitch(twitchUser.ID)
	if err != nil {
		log.Println(err)
		if err == gorm.ErrRecordNotFound {
			msg := fmt.Sprintf("%s usuário não encontrado. Dê !join para se cadastrar", twitchUser.DisplayName)
			return msg, nil
		}
		msg := fmt.Sprintf("%s não foi possível assinar sua presença", twitchUser.DisplayName)
		return msg, err
	}

	check, err := s.CheckPresentToday(customer.UUID)
	if err != nil {
		log.Println("Erro ao verificar presença:", err)
		return "", err
	}

	if check {
		msg := fmt.Sprintf("%s você já assinou presença hoje!", twitchUser.DisplayName)
		return msg, nil
	}

	if err := s.addPresencaCubes(customer); err != nil {
		msg := fmt.Sprintf("%s não foi possível assinar sua presença", twitchUser.DisplayName)
		return msg, err
	}

	s.addStreakCubes(*customer.IdTwitch)

	msg := fmt.Sprintf("%s presença assinada com sucesso!", twitchUser.DisplayName)
	return msg, nil
}

func (s *PointsService) CubesToDatapoints(twitchUser twitch.User) (string, error) {

	loyaltyUser, err := s.loyaltyRepository.GetCustomerByTwitch(twitchUser.ID)
	if err != nil {
		return "", err
	}

	qtde := int(loyaltyUser.NrPoints / 1000)

	if qtde >= 1 {
		products := []models.ProductPoints{}
		for i := 1; i <= qtde; i++ {
			products = append(products, models.NewTroca())
		}

		if err := s.streamElementsRepository.AddPoints(twitchUser.Name, int64(qtde*100)); err != nil {
			log.Println(err)
			msg := fmt.Sprintf("%s não foi possível realizar a troca", twitchUser.DisplayName)
			return msg, err
		}

		if err := s.loyaltyRepository.AddPoints(loyaltyUser.UUID, products); err != nil {
			log.Println(err)
			msg := fmt.Sprintf("%s não foi possível realizar a troca", twitchUser.DisplayName)
			return msg, err
		}

		msg := fmt.Sprintf("%s troca realizada com sucesso: %d cubos trocados!", twitchUser.DisplayName, qtde*1000)
		return msg, nil
	}

	msg := fmt.Sprintf("%s você não tem cubos suficientes. Junte 1.000 cubos!", twitchUser.DisplayName)
	return msg, nil

}
