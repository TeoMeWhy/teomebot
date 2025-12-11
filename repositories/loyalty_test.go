package repositories

import (
	"teomebot/config"
	"teomebot/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoyalCreateUser(t *testing.T) {

	settings := &config.Config{
		LoyaltyServiceURI: "http://localhost:8081",
	}

	loyaltyRepo := NewLoyaltyRepository(settings)

	uuid, err := loyaltyRepo.CreateCustomerByTwitch("test_user_id_01")
	assert.NoError(t, err)
	assert.NotEmpty(t, uuid)

	err = loyaltyRepo.DeleteCustomer(uuid)
	assert.NoError(t, err)

}

func TestLoyalAddPoints(t *testing.T) {

	settings := &config.Config{
		LoyaltyServiceURI: "http://localhost:8081",
	}

	loyaltyRepo := NewLoyaltyRepository(settings)

	uuid, err := loyaltyRepo.CreateCustomerByTwitch("test_user_id_01")
	assert.NoError(t, err)
	assert.NotEmpty(t, uuid)

	chat := models.NewChatMessage()

	products := []models.ProductPoints{chat}

	err = loyaltyRepo.AddPoints(uuid, products)
	assert.NoError(t, err)

	customer, err := loyaltyRepo.GetCustomer(uuid)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), customer.NrPoints)

	err = loyaltyRepo.DeleteCustomer(uuid)
	assert.NoError(t, err)

}
