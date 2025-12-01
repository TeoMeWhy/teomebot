package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Customer struct {
	UUID             string    `json:"uuid" gorm:"primaryKey"`
	DescCustomerName string    `json:"customer_name"`
	CodCPF           *string   `json:"cpf" gorm:"unique"`
	DescEmail        string    `json:"email" gorm:"unique"`
	IdTwitch         *string   `json:"twitch" gorm:"unique"`
	IdYouTube        *string   `json:"youtube" gorm:"unique"`
	IdBlueSky        *string   `json:"bluesky" gorm:"unique"`
	IdInstagram      *string   `json:"instagram" gorm:"unique"`
	NrPoints         int64     `json:"points"`
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
}

type CreateUserResp struct {
	Customer Customer `json:"customer"`
	Status   string   `json:"status"`
}

func CreateUser(twitchID string) (string, error) {
	user := map[string]string{
		"twitch": twitchID,
	}

	bodyUser, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	url := "http://points:8081/customers/"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyUser))
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", errors.New("erro na criação do usuário")
	}

	respContent, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	userResp := &CreateUserResp{}
	if err := json.Unmarshal(respContent, &userResp); err != nil {
		return "", err
	}

	return userResp.Customer.UUID, nil
}

func UpdateUser(user map[string]string) error {

	bodyUser, err := json.Marshal(user)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("http://points:8081/customers/%s", user["uuid"])
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyUser))
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("erro na criação do usuário")
	}

	return nil
}

func GetCustomer(twitchID string) (*Customer, error) {

	url := "http://points:8081/customers/?twitch=%s"
	url = fmt.Sprintf(url, twitchID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro na requisição. statuscode: %d", resp.StatusCode)
	}

	bodyReader, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	customers := []Customer{}
	if err := json.Unmarshal(bodyReader, &customers); err != nil {
		return nil, err
	}

	customer := customers[0]
	return &customer, nil
}
