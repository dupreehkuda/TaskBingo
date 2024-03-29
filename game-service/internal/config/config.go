package config

import (
	"flag"

	"github.com/caarlos0/env/v6"
)

// Config provides service address and paths to database
type Config struct {
	Address            string `env:"GAME_SERVICE_ADDRESS" envDefault:":8082"`
	UserServiceAddress string `env:"FULL_USER_SERVICE_ADDRESS"`
	CurrentDomain      string `env:"CURRENT_DOMAIN"`
}

// New creates new Config
func New() *Config {
	var config = Config{}
	var err = env.Parse(&config)
	if err != nil {
		panic(err)
	}

	flag.StringVar(&config.Address, "a", config.Address, "Launch address")
	flag.StringVar(&config.UserServiceAddress, "u", config.UserServiceAddress, "User service address")
	flag.StringVar(&config.CurrentDomain, "d", config.CurrentDomain, "Current domain of the app")
	flag.Parse()

	return &config
}
