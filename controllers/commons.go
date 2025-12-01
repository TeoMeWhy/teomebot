package controllers

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"teomebot/clients/streamelements"
	"teomebot/utils"

	"github.com/gempir/go-twitch-irc/v4"
)

var conDB, _ = utils.OpenDBConnection()

var clientStreamElements = streamelements.NewStreamElementsClient()

func MegaController(u twitch.User) string {
	numbers := numbersGenerate()
	txt := strings.Join(numbers, "-")
	txt = fmt.Sprintf("%s sua mega: %s", u.Name, txt)
	return txt
}

func numbersGenerate() []string {
	values := []int{}
	for len(values) < 6 {
		value := rand.Intn(60) + 1
		if isInList(value, values) {
			continue
		}

		values = append(values, value)
	}

	sort.Ints(values)

	numbers := []string{}
	for _, v := range values {
		numbers = append(numbers, fmt.Sprintf("%d", v))
	}

	return numbers
}

func isInList(n int, list []int) bool {
	for _, i := range list {
		if i == n {
			return true
		}
	}
	return false
}
