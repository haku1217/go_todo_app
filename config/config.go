package config

import "github.com/caarlos0/env/v6"

var envParse = env.Parse

type Config struct {
	Env  string `env:"TODO_ENV" envDefault:"dev"`
	Port int    `env:"PORT" envDefault:"80"`
}

func New() (*Config, error) {
	cfg := &Config{}

	if err := envParse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
