package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"teomebot/config"
)

type RetroRepository struct {
	URI        string
	HttpClient *http.Client
}

type retroResponse struct {
	Report string `json:"report"`
}

func (r *RetroRepository) GetUserRetro(uuid, nick string) (*string, error) {

	log.Println("AQUI PRE MERDA:", uuid, nick, r.URI)

	url := fmt.Sprintf("%s/retro?id=%s&name=%s?source=twitch", r.URI, uuid, nick)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("user not found")
	} else if resp.StatusCode != http.StatusOK {
		log.Println(string(bodyBytes))
		return nil, errors.New("erro desconhecido")
	}

	payload := &retroResponse{}
	if err := json.Unmarshal(bodyBytes, &payload); err != nil {
		return nil, err
	}

	return &payload.Report, nil
}

func NewRetroRepository(settings *config.Config) *RetroRepository {
	return &RetroRepository{
		URI:        settings.RetroServiceURI,
		HttpClient: &http.Client{},
	}
}
