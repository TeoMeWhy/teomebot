package repositories

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"teomebot/config"
)

var errUnexpectedStatusCode = fmt.Errorf("unexpected status code")
var errUnexpectedResponse = fmt.Errorf("unexpected response")

type QueryPayloadRequest struct {
	Query string `json:"query"`
}

type QueryPayloadResponse struct {
	Response string `json:"response"`
	Error    string `json:"error"`
}

type RagiaClient struct {
	url        string
	httpClient *http.Client
}

func (c *RagiaClient) GetQueryResponse(query string) (string, error) {

	payloadRequestBytes, err := json.Marshal(&QueryPayloadRequest{Query: query})
	if err != nil {
		return "", err
	}

	bodyRequest := bytes.NewBuffer(payloadRequestBytes)

	resp, err := c.httpClient.Post(c.url, "application/json", bodyRequest)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errUnexpectedStatusCode
	}

	var payloadResponse QueryPayloadResponse
	err = json.NewDecoder(resp.Body).Decode(&payloadResponse)
	if err != nil {
		return "", err
	}

	if payloadResponse.Error != "" {
		return "", fmt.Errorf(payloadResponse.Error)
	}

	if payloadResponse.Response == "" {
		return "", errUnexpectedResponse
	}

	return payloadResponse.Response, nil
}

func NewRagiaClient(cfg *config.Config) *RagiaClient {

	return &RagiaClient{
		url:        cfg.RagiaURL,
		httpClient: &http.Client{},
	}
}
