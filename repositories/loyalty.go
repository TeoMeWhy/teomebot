package repositories

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"teomebot/config"
	"teomebot/models"
	"time"
)

type Product struct {
	ProductID   string `json:"product_id"`
	ProductQtde int64  `json:"product_qtd"`
	Points      int64  `json:"points"`
}

type Transacation struct {
	CustomerID   string    `json:"customer_id"`
	Points       int64     `json:"points"`
	SystemOrigin string    `json:"system_origin"`
	Products     []Product `json:"products"`
}

type Customer struct {
	UUID             string    `json:"uuid"`
	DescCustomerName string    `json:"customer_name"`
	CodCPF           *string   `json:"cpf"`
	DescEmail        string    `json:"email"`
	IdTwitch         *string   `json:"twitch"`
	IdYouTube        *string   `json:"youtube"`
	IdBlueSky        *string   `json:"bluesky"`
	IdInstagram      *string   `json:"instagram"`
	NrPoints         int64     `json:"points"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type CreateUserResp struct {
	Customer Customer `json:"customer"`
	Status   string   `json:"status"`
}

type LoyaltyRepository struct {
	URI        string
	HttpClient *http.Client
}

func (r *LoyaltyRepository) GetCustomerByTwitch(twitchID string) (*Customer, error) {

	url := "%s/customers/?twitch=%s"
	url = fmt.Sprintf(url, r.URI, twitchID)

	resp, err := r.HttpClient.Get(url)
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

func (r *LoyaltyRepository) CreateCustomerByTwitch(twitchID string) (string, error) {

	user := map[string]string{
		"twitch": twitchID,
	}

	bodyUser, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/customers/", r.URI)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyUser))
	if err != nil {
		return "", err
	}

	resp, err := r.HttpClient.Do(req)
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

func (r *LoyaltyRepository) UpdateCustomer(customer map[string]string) error {

	bodyUser, err := json.Marshal(customer)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/customers/%s", r.URI, customer["uuid"])
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyUser))
	if err != nil {
		return err
	}

	resp, err := r.HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("erro na criação do usuário")
	}

	return nil
}

func (r *LoyaltyRepository) AddPoints(userId string, products []models.ProductPoints) error {

	tProducts := []Product{}
	var total int64
	for _, p := range products {
		tProducts = append(tProducts, Product{
			ProductID:   p.GetCod(),
			ProductQtde: p.GetQtde(),
			Points:      p.GetValue(),
		})
		total += p.GetValue()
	}

	transaction := Transacation{
		CustomerID:   userId,
		Points:       total,
		SystemOrigin: "twitch",
		Products:     tProducts,
	}

	url := fmt.Sprintf("%s/transactions", r.URI)
	bodyReq, err := json.Marshal(&transaction)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyReq))
	if err != nil {
		return err
	}

	client := r.HttpClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("erro ao criar transaction %d", resp.StatusCode)
	}

	return nil

}

func (r *LoyaltyRepository) DeleteCustomer(customerID string) error {

	url := fmt.Sprintf("%s/customers/%s", r.URI, customerID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	resp, err := r.HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("erro ao deletar customer %d", resp.StatusCode)
	}

	return nil

}

func NewLoyaltyRepository(settings *config.Config) *LoyaltyRepository {
	return &LoyaltyRepository{
		URI:        settings.LoyaltyServiceURI,
		HttpClient: &http.Client{},
	}
}
