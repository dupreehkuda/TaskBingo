package config

import (
	"flag"
	"strings"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Address          string   `env:"GAME_SERVICE_ADDRESS" envDefault:":8082"`
	DatabaseDSN      string   `env:"DATABASE_DSN"`
	JWTSecret        string   `env:"secret"`
	CurrentDomain    string   `env:"CURRENT_DOMAIN" envDefault:"localhost"`
	AllowedWSOrigins []string `env:"ALLOWED_WS_ORIGINS" envDefault:"http://localhost:3000,https://taskbingo.com" envSeparator:","`
	MigrationsPath   string   `env:"MIGRATIONS_PATH" envDefault:"./migrations/postgresql"`
}

func New() *Config {
	var c Config
	if err := env.Parse(&c); err != nil {
		panic(err)
	}

	flag.StringVar(&c.Address, "a", c.Address, "Listen address")
	flag.StringVar(&c.DatabaseDSN, "d", c.DatabaseDSN, "Postgres DSN")
	flag.StringVar(&c.CurrentDomain, "domain", c.CurrentDomain, "Current domain")
	flag.Parse()

	for i, o := range c.AllowedWSOrigins {
		c.AllowedWSOrigins[i] = strings.TrimSpace(o)
	}
	return &c
}

// IsLocal reports whether the server is running in the local-dev profile.
func (c *Config) IsLocal() bool { return c.CurrentDomain == "localhost" }
