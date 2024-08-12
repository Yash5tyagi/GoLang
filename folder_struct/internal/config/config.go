package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB     *DB
	Server *Server
}

type DB struct {
	DBDriver   string
	DBDatabase string
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
}

type Server struct {
	Port string
	Env  string
}

// LoadConfig loads the configuration values from environment variables or a config file
func LoadConfig() (*Config, error) {
	env, err := LoadENV()
	if err != nil {
		log.Println("error loading env")
		return nil, err
	}

	// Initialize a new Config struct
	config := &Config{
		DB: &DB{
			DBDriver:   env.DBDriver,
			DBDatabase: env.DBDatabase,
			DBUsername: env.DBUsername,
			DBPassword: env.DBPassword,
			DBHost:     env.DBHost,
			DBPort:     env.DBPort,
		},
		Server: &Server{
			Port: env.ServerPort,
			Env:  env.Environment,
		},
	}

	// Read configuration values from environment variables
	viper.AutomaticEnv()

	// Read configuration values from the config file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./.config")

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found. Using environment variables.")
		} else {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	// Load configuration values into the Config struct
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	log.Println(config)
	return config, nil
}

type CtxKey struct {
	RequestID string `json:"request_id"`
	Session   string `json:"session"`
}
