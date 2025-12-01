package controllers

import (
	"fmt"
	"log"
	"teomebot/clients/streamelements"
	"teomebot/models"
	"teomebot/services"

	"github.com/gempir/go-twitch-irc/v4"
)

func MessageChatController(u twitch.User) {

	user := &models.TwitchUser{}
	res := conDB.First(&user, "twitch_id = ?", u.ID)
	if res.Error != nil {
		log.Println(u.Name, res.Error)
		return
	}

	chatProduct := models.NewChatMessage()
	if err := services.AddPoints(user.UUID, chatProduct); err != nil {
		log.Println(u.Name, user.UUID, err)
	}
}

func PresentController(u twitch.User) string {

	user := &models.TwitchUser{}
	res := conDB.First(&user, "twitch_id = ?", u.ID)
	if res.Error != nil {
		log.Println(res.Error)
		txt := fmt.Sprintf("%s usuário não encontrado. Dê !join", u.Name)
		return txt
	}

	if models.PresentExists(user.UUID) {
		txt := fmt.Sprintf("%s lista já foi assinada. Pare!", u.Name)
		return txt
	}

	presencaPoints := models.NewPresent()
	if err := services.AddPoints(user.UUID, presencaPoints); err != nil {
		log.Println(err)
		return ""
	}

	present := models.NewPresentUser(user.UUID)
	if res := conDB.Create(present); res.Error != nil {
		log.Println(res.Error)
		return ""
	}

	// Atualiza a tabela de streak
	streak := models.LoadStreak(user.UUID)
	streak.Qtd++
	conDB.Save(streak)
	if streak.Qtd%5 == 0 {

		for i := int(streak.Qtd) / 5; i >= 1; i-- {

			p := models.NewStreakPresent()
			err := services.AddPoints(user.UUID, p)

			if err != nil {
				log.Println(err)
				txt := fmt.Sprintf("%s tentativa falha de atribuir pontos de streak", u.Name)
				return txt
			}
		}
	}

	txt := fmt.Sprintf("%s lista assinada com sucesso.", u.Name)
	return txt
}

func ProfileController(u twitch.User) string {

	user := &models.TwitchUser{}

	if res := conDB.First(&user, "twitch_id = ?", u.ID); res.Error != nil {
		return fmt.Sprintf("%s usuário não encontrado. Dê !join para participar.", user.TwitchNick)
	}

	proba, err := services.GetUserChurnProb(user.UUID)
	if err != nil {

		if err.Error() == "user not found" {
			return fmt.Sprintf("%s usuário não encontrado. Dê !join ou volte amanhã.", user.TwitchNick)
		}

		log.Println(err)
		return fmt.Sprintf("%s erro ao obter seu churn score.", user.TwitchNick)
	}

	if models.ProfileUserExists(user.UUID) {
		return fmt.Sprintf("%s Proba. Churn: %.2f%%", user.TwitchNick, proba*100)
	}

	profile := models.NewProfileUser(user.UUID)
	if res := conDB.Save(profile); res.Error != nil {
		return fmt.Sprintf("%s erro ao salvar seu profile de hoje", user.TwitchNick)
	}

	p := models.NewChurn(proba)
	if err := services.AddPoints(user.UUID, p); err != nil {
		return fmt.Sprintf("%s deu ruim em adicionar seus pontos.", user.TwitchNick)
	}

	return fmt.Sprintf("%s Proba. Churn: %.2f%% | +%d cubos", user.TwitchNick, proba*100, p.VlProduct)

}

func GetUserPointsController(u twitch.User) string {

	customer, err := services.GetCustomer(u.ID)
	if err != nil {
		log.Println(err)
		return fmt.Sprintf("%s não foi possível obter seus cubos", u.Name)
	}

	return fmt.Sprintf("%s você tem %d cubos", u.Name, customer.NrPoints)
}

func TrocaController(u twitch.User) string {

	customer, err := services.GetCustomer(u.ID)
	if err != nil {
		return fmt.Sprintf("%s não encontramos seu usuário. Dê !join", u.DisplayName)
	}

	if customer.NrPoints < 1000 {
		return fmt.Sprintf("%s cubos insuficientes para troca. Junte 1000 cubos.", u.DisplayName)
	}

	streamReq := &streamelements.AddPointsRequest{UserName: u.Name, Amount: 100}
	if err := clientStreamElements.AddPoints(streamReq); err != nil {
		log.Println(err)
		return fmt.Sprintf("%s erro ao atribuir seus datapoints no streamElements", u.DisplayName)
	}

	prod := models.NewTroca()
	if err := services.AddPoints(customer.UUID, prod); err != nil {
		streamReq.Amount = -100
		clientStreamElements.AddPoints(streamReq)
		return fmt.Sprintf("%s erro ao realizar sua troca", u.DisplayName)
	}

	return fmt.Sprintf("%s troca realizada com sucesso", u.DisplayName)
}
