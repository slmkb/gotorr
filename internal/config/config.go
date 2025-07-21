package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL       string `json:"db_url"`
	CurrentUser string `json:"current_user_name"`
}

func Read() Config {
	ConfigFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}
	}
	dat, err := os.ReadFile(ConfigFilePath)
	if err != nil {
		panic(err)
	}

	var cfg Config
	if err := json.Unmarshal(dat, &cfg); err != nil {
		panic(err)
	}
	return cfg
}

func (cfg *Config) SetUser(userName string) error {
	if len(userName) == 0 {
		return errors.New("username can't be empty")
	}

	cfg.CurrentUser = userName
	if err := write(*cfg); err != nil {
		return fmt.Errorf("write data error: %w", err)
	}
	return nil
}

func write(cfg Config) error {
	dat, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("marshal data error: %w", err)
	}
	ConfigFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	os.WriteFile(ConfigFilePath, dat, 0600)
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, configFileName), nil
}
