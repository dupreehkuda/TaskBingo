package config

import (
	"flag"

	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"
)

// Config provides service address and paths to database
type Config struct {
	Address            string `env:"GAME_SERVICE_ADDRESS" envDefault:":8082"`
	UserServiceAddress string `env:"FULL_USER_SERVICE_ADDRESS"`
	TaskServiceAddress string `env:"FULL_TASK_SERVICE_ADDRESS"`
}

// New creates new Config
func New(logger *zap.Logger) *Config {
	var config = Config{}
	var err = env.Parse(&config)
	if err != nil {
		logger.Error("Error occurred when parsing config", zap.Error(err))
	}

	flag.StringVar(&config.Address, "a", config.Address, "Launch address")
	flag.StringVar(&config.UserServiceAddress, "u", config.UserServiceAddress, "User service address")
	flag.StringVar(&config.TaskServiceAddress, "t", config.TaskServiceAddress, "Task service address")
	flag.Parse()

	return &config
}
