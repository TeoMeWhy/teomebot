package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"teomebot/models"
)

func AddPoints(userId string, p models.ProductPoints) error {

	transaction := map[string]any{
		"customer_id":   userId,
		"points":        p.GetValue(),
		"system_origin": "twitch",
		"products": []map[string]any{
			{
				"product_id":  p.GetCod(),
				"product_qtd": p.GetQtde(),
				"points":      p.GetValue(),
			},
		},
	}

	url := "http://points:8081/transactions"
	bodyReq, err := json.Marshal(&transaction)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyReq))
	if err != nil {
		return err
	}

	client := &http.Client{}
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
