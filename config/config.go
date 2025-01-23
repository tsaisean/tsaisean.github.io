package config

import (
	"github.com/pelletier/go-toml/v2"
	"os"
)

type Config struct {
	AppName  string   `toml:"appName"`
	Env      string   `toml:"env"`
	Database Database `toml:"database"`
}

type Database struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
}

func LoadConfig() (*Config, error) {
	return loadConfig("config/config.toml")
}

func loadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := toml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
