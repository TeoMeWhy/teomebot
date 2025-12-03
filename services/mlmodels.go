package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type userProba struct {
	UUID  string  `json:"uuid"`
	Proba float64 `json:"prob"`
}

func GetUserChurnProb(userId string) (float64, error) {

	url := "http://models:8080/churn_score/%s"
	url = fmt.Sprintf(url, userId)

	resp, err := http.Get(url)
	if err != nil {
		return 0., err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0., err
	}

	if resp.StatusCode == http.StatusNotFound {
		return 0., errors.New("user not found")
	} else if resp.StatusCode != http.StatusOK {
		log.Println(string(bodyBytes))
		return 0., errors.New("erro desconhecido")
	}

	userproba := &userProba{}
	if err := json.Unmarshal(bodyBytes, &userproba); err != nil {
		return 0., err
	}

	return userproba.Proba, nil
}
