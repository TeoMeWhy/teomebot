package services

import (
	"fmt"
	"log"
	"teomebot/config"
	"teomebot/errors"
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

	user, err := s.userRepository.GetUserByField("twitch_id", twitchUser.ID)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Println(err)
		}
		return
	}

	product := models.NewChatMessage()
	products := []models.ProductPoints{product}

	if err := s.loyaltyRepository.AddPoints(user.UUID, products); err != nil {
		log.Println(err)
		return
	}

}

func (s *PointsService) addPresencaCubes(user *repositories.TwitchUser) error {

	present, err := s.presentRepository.LoadLastPresent(user)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err)
		return err
	}

	if present.CreatedAt.Format("2006-01-02") == time.Now().Format("2006-01-02") {
		return errors.ErrPresencaAssinadaAnterior
	}

	newPresent, err := s.presentRepository.CreatePresenca(user)
	if err != nil {
		log.Println(err)
		return err
	}

	products := []models.ProductPoints{models.NewPresent()}
	if err := s.loyaltyRepository.AddPoints(user.UUID, products); err != nil {
		s.presentRepository.DeletePresenca(newPresent)
		log.Println(err)
		return err
	}

	return nil
}

func (s *PointsService) addStreakCubes(user *repositories.TwitchUser) {

	streak, err := s.presentRepository.LoadLastUpdatedStreak(user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := s.presentRepository.CreateStreak(user); err != nil {
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
			if err := s.loyaltyRepository.AddPoints(user.UUID, products); err != nil {
				log.Println(err)
				return
			}
		}
		return
	}

	if err := s.presentRepository.CreateStreak(user); err != nil {
		log.Println(err)
		return
	}

}

func (s *PointsService) MgmtPresenca(twitchUser twitch.User) (string, error) {

	user, err := s.userRepository.GetUserByField("twitch_id", twitchUser.ID)
	if err != nil {
		log.Println(err)
		if err == gorm.ErrRecordNotFound {
			msg := fmt.Sprintf("%s usuário não encontrado. Dê !join para se cadastrar", twitchUser.DisplayName)
			return msg, nil
		}
		msg := fmt.Sprintf("%s não foi possível assinar sua presença", twitchUser.DisplayName)
		return msg, err
	}

	if err := s.addPresencaCubes(user); err != nil {

		if err == errors.ErrPresencaAssinadaAnterior {
			msg := fmt.Sprintf("%s você já assinou presença hoje!", twitchUser.DisplayName)
			return msg, nil
		}

		msg := fmt.Sprintf("%s não foi possível assinar sua presença", twitchUser.DisplayName)
		return msg, err
	}

	s.addStreakCubes(user)

	msg := fmt.Sprintf("%s presença assinada com sucesso!", twitchUser.DisplayName)
	return msg, nil
}
