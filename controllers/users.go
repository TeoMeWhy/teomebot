package controllers

import (
	"fmt"
	"log"
	"teomebot/models"
	"teomebot/services"

	"github.com/gempir/go-twitch-irc/v4"
)

func ExecCreateOrUpdateUser(twitchUser *twitch.User) error {

	user, err := models.GetUserByField("twitch_id", twitchUser.ID, conDB)
	if err != nil {

		user, err = models.GetUserByField("twitch_nick", twitchUser.Name, conDB)
		if err != nil {

			var userID string

			customer, err := services.GetCustomer(twitchUser.ID)
			if err != nil {
				userID, err = services.CreateUser(twitchUser.ID)
				if err != nil {
					log.Println(err)
					return err
				}
			} else {
				userID = customer.UUID
			}

			user = &models.TwitchUser{
				UUID:       userID,
				TwitchId:   twitchUser.ID,
				TwitchNick: twitchUser.Name,
			}

			return user.Create(conDB)
		}
	}

	if user.TwitchId == twitchUser.ID && user.TwitchNick == twitchUser.Name {
		return fmt.Errorf("usuário sem atualização")
	}

	user.TwitchId = twitchUser.ID
	user.TwitchNick = twitchUser.Name
	if err := user.Update(conDB); err != nil {
		return err
	}

	userMap := map[string]string{
		"uuid":   user.UUID,
		"twitch": user.TwitchId,
	}

	return services.UpdateUser(userMap)
}

func RetroController(u twitch.User) string {

	user := &models.TwitchUser{}

	if res := conDB.First(&user, "twitch_id = ?", u.ID); res.Error != nil {
		return fmt.Sprintf("%s usuário não encontrado. Dê !join para participar.", user.TwitchNick)
	}

	retro, err := services.GetUserRetro(user.UUID)
	if err != nil {

		if err.Error() == "user not found" {
			return fmt.Sprintf("%s usuário não encontrado. Dê !join ou volte amanhã.", user.TwitchNick)
		}

		log.Println(err)
		return fmt.Sprintf("%s erro ao obter a sua retro.", user.TwitchNick)
	}

	txt := fmt.Sprintf("%s %s", u.DisplayName, *retro)

	return txt
}
