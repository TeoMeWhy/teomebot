package streamelements

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type StreamElementsClient struct {
	URL     string `json:"url"`
	Channel string `json:"channel"`
	Token   string `json:"token"`
}

func (c *StreamElementsClient) MakeHeader() http.Header {

	header := http.Header{}
	header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	header.Add("Content-Type", "application/json")
	header.Add("Accept", "")

	return header
}

func (c *StreamElementsClient) AddPoints(req *AddPointsRequest) error {
	url := fmt.Sprintf("%s/points/%s/%s/%d", c.URL, c.Channel, req.UserName, req.Amount)

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
		return fmt.Errorf("erro: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Println(string(body))

	return nil

}

func NewStreamElementsClient() *StreamElementsClient {

	// if err := godotenv.Load(".env"); err != nil {
	// 	panic(err)
	// }

	return &StreamElementsClient{
		URL:     "https://api.streamelements.com/kappa/v2",
		Channel: os.Getenv("STREAMELEMENTS_ACCOUNT_ID"),
		Token:   os.Getenv("STREAMELEMENTS_TOKEN"),
	}

}
