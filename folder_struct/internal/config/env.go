package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Env struct {
	DBDriver    string `env:"DB_DRIVER"`
	DBDatabase  string `env:"DB_DATABASE"`
	DBUsername  string `env:"DB_USERNAME"`
	DBPassword  string `env:"DB_PASSWORD"`
	DBHost      string `env:"DB_HOST"`
	DBPort      string `env:"DB_PORT"`
	ServerPort  string `env:"SERVER_PORT"`
	Environment string `env:"ENVIRONMENT"`
	BaseURL     string `env:"BASEURL"`
}

func LoadENV() (cfg *Env, err error) {
	godotenv.Load("./.secrets/.env")
	cfg = new(Env)
	err = env.Parse(cfg)
	if err != nil {
		return
	}
	return
}
