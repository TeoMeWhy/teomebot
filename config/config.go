package config

import "github.com/caarlos0/env"

type Config struct {
	LoyaltyServiceURI string `env:"LOYALTY_SERVICE_URI" envDefault:"http://localhost:8081"`
	RetroServiceURI   string `env:"RETRO_SERVICE_URI" envDefault:"http://localhost:8082"`

	StreamElementsURI     string `env:"STREAM_ELEMENTS_URI" envDefault:"https://api.streamelements.com/kappa/v2"`
	StreamElementsChannel string `env:"STREAM_ELEMENTS_CHANNEL"`
	StreamElementsToken   string `env:"STREAM_ELEMENTS_TOKEN"`

	TwitchChannel  string `env:"TWITCH_CHANNEL"`
	TwitchBot      string `env:"TWITCH_BOT"`
	TwitchOauthBot string `env:"TWITCH_OAUTH_BOT"`

	DsnMysql string `env:"DSN_MYSQL"`
}

func LoadConfig() (*Config, error) {

	config := Config{}
	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
