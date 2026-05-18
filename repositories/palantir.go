package repositories

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"teomebot/config"
)

type PalantirRepository struct {
	uri        string
	httpClient *http.Client
}

type PredictionsClassification map[string]float64

type PredictionResponse struct {
	Predictions map[string]PredictionsClassification `json:"predictions"`
	Err         *string                              `json:"error,omitempty"`
}

func (c *PalantirRepository) GetPrediction(modelName, id string) (PredictionsClassification, error) {

	payload := map[string]interface{}{
		"model_name": modelName,
		"id":         id,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling payload: %v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", c.uri+"/predict", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("Error making request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	var predictionResponse PredictionResponse
	if err := json.NewDecoder(resp.Body).Decode(&predictionResponse); err != nil {
		log.Printf("Error decoding response: %v", err)
		return nil, err
	}

	if predictionResponse.Err != nil {
		log.Printf("Error from prediction service: %s", *predictionResponse.Err)
		return nil, err
	}

	prediction, ok := predictionResponse.Predictions[id]
	if !ok {
		log.Printf("Prediction not found for id: %s", id)
		return nil, err
	}

	return prediction, nil
}

func NewPalantirRepository(settings *config.Config, httpClient *http.Client) *PalantirRepository {
	return &PalantirRepository{
		uri:        settings.PalantirURI,
		httpClient: httpClient,
	}
}
