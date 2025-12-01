package repositories

import (
	"teomebot/config"
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
