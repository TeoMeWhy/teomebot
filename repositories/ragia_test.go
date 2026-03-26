package repositories

import (
	"net/http"
	"teomebot/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name             string
	Query            string
	QueryNotExpected string
	Error            error
}

type TestStructWithError struct {
	Name          string
	Query         string
	QueryExpected string
	Error         error
}

func TestRagiaNoError(t *testing.T) {

	settings := &config.Config{
		RagiaURL: "http://192.168.0.18:5003/predict",
	}

	client := &RagiaClient{
		url:        settings.RagiaURL,
		httpClient: &http.Client{},
	}

	tests := []TestStruct{
		{
			Name:             "Test 1 - Como começar em dados?",
			Query:            "Como começar em dados?",
			QueryNotExpected: "",
			Error:            nil,
		},
		{
			Name:             "Test 2 - Qual o horário de live?",
			Query:            "Qual o horário de live?",
			QueryNotExpected: "",
			Error:            nil,
		},
	}

	for _, tt := range tests {

		t.Run(tt.Name, func(t *testing.T) {

			response, err := client.GetQueryResponse(tt.Query)

			assert.NotEqual(t, tt.QueryNotExpected, response)
			assert.NoError(t, err)

		})

	}
}

func TestRagiaError(t *testing.T) {

	settings := &config.Config{
		RagiaURL: "http://192.168.0.18:5003/predict",
	}

	client := &RagiaClient{
		url:        settings.RagiaURL,
		httpClient: &http.Client{},
	}

	tests := []TestStructWithError{
		{
			Name:          "Test 1 - ''",
			Query:         "",
			QueryExpected: "",
			Error:         errUnexpectedStatusCode,
		},
		{
			Name:          "Test 2 - Como fazer uma macarronada?",
			Query:         "Como fazer uma macarronada?",
			QueryExpected: "",
			Error:         errUnexpectedResponse,
		},
		{
			Name:          "Test 3 - Que dia lindo",
			Query:         "Que dia lindo",
			QueryExpected: "",
			Error:         errUnexpectedStatusCode,
		},
	}

	for _, tt := range tests {

		t.Run(tt.Name, func(t *testing.T) {

			response, err := client.GetQueryResponse(tt.Query)
			assert.Equal(t, tt.QueryExpected, response)
			assert.NotNil(t, err)
		})

	}
}
