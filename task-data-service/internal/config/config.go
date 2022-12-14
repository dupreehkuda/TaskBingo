package config

import (
	"flag"

	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"
)

// Config provides service address and paths to database
type Config struct {
	Address      string `env:"RUN_ADDRESS" envDefault:":8083"`
	DatabasePath string `env:"REDIS_ADDR"`
	DatabasePass string `env:"REDIS_PASS"`
}

// New creates new Config
func New(logger *zap.Logger) *Config {
	var config = Config{}
	var err = env.Parse(&config)
	if err != nil {
		logger.Error("Error occurred when parsing config", zap.Error(err))
	}

	flag.StringVar(&config.Address, "a", config.Address, "Launch address")
	flag.StringVar(&config.DatabasePath, "d", config.DatabasePath, "Path to database")
	flag.StringVar(&config.DatabasePath, "p", config.DatabasePath, "Path to database")
	flag.Parse()

	return &config
}
