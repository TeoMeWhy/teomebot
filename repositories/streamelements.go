package repositories

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"teomebot/config"
)

type StreamElementsClient struct {
	URI     string
	Channel string
	Token   string
}

func (c *StreamElementsClient) MakeHeader() http.Header {
	header := http.Header{}
	header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	header.Add("Content-Type", "application/json")
	header.Add("Accept", "")
	return header
}

func (c *StreamElementsClient) AddPoints(username string, amount int64) error {
	url := fmt.Sprintf("%s/points/%s/%s/%d", c.URI, c.Channel, username, amount)

	request, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}

	request.Header = c.MakeHeader()

	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		log.Println(string(body))

		return fmt.Errorf("erro: status code %d", resp.StatusCode)
	}

	return nil

}

func NewStreamElementsClient(settings *config.Config) *StreamElementsClient {
	return &StreamElementsClient{
		URI:     settings.StreamElementsURI,
		Channel: settings.StreamElementsChannel,
		Token:   settings.StreamElementsToken,
	}
}
