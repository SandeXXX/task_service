package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type Config struct {
	Env        string     `json:"env" env-default:"local"`
	Database   Database   `json:"database"`
	AppName    string     `json:"app_name"`
	HTTPServer HTTPServer `json:"http_server"`
}

type HTTPServer struct {
	Port int `json:"port" env:"HTTP_SERVER_PORT"`
}

type Database struct {
	Host     string `json:"host" env:"POSTGRES_HOST"`
	Port     int    `json:"port" env:"POSTGRES_PORT"`
	Database string `json:"database" env:"POSTGRES_DBNAME"`
	User     string `json:"user" env:"POSTGRES_USER"`
	Pass     string `json:"pass" env:"POSTGRES_PASSWORD"`
}

func NewConfig() (*Config, error) {
	var config Config
	if err := config.readConfig(); err != nil {
		return nil, err
	}

	return &config, nil
}

func (config *Config) readConfig() error {
	configPath := os.Getenv("CONFIG_PATH")
	if strings.TrimSpace(configPath) == "" {
		return errors.New("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("config file does not exist - `%s`", configPath))
	}

	err := cleanenv.ReadConfig(configPath, config)
	if err != nil {
		return err
	}
	return nil
}
