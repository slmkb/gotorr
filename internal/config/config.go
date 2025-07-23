package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gotorrconfig.json"

type Config struct {
	DbURL       string `json:"db_url"`
	CurrentUser string `json:"current_user_name"`
}

func Read() (Config, error) {
	ConfigFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	dat, err := os.ReadFile(ConfigFilePath)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := json.Unmarshal(dat, &cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUser = userName
	return write(*cfg)
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
	return filepath.Join(homeDir, configFileName), nil
}
