package config

type Config struct {
	LoyaltyServiceURI     string
	RetroServiceURI       string
	StreamElementsURI     string `default:"https://api.streamelements.com/kappa/v2"`
	StreamElementsChannel string
	StreamElementsToken   string
}
