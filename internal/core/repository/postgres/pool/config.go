package core_postgres_pool

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host     string        `envconfig:"HOST" default:"localhost" required:"true"`
	Port     int           `envconfig:"PORT" default:"5432"`
	User     string        `envconfig:"USER" default:"postgres" required:"true"`
	Password string        `envconfig:"PASSWORD" required:"true"`
	DBName   string        `envconfig:"DB" default:"postgres" required:"true"`
	Timeout  time.Duration `envconfig:"TIMEOUT" default:"30s" required:"true"`
}

func NewConfig() (Config, error) {
	var cfg Config

	if err := envconfig.Process("POSTGRES", &cfg); err != nil {
		return Config{}, fmt.Errorf("process env vars: %w", err)
	}

	return cfg, nil
}

func NewConfigMust() Config {
	cfg, err := NewConfig()
	if err != nil {
		panic(fmt.Errorf("get postgres connection pool config: %w", err))
	}

	return cfg
}
