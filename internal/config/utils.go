package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL       string `json:"db_url"`
	CurrentUser string `json:"current_user_name"`
}

func Read() (*Config, error) {
	var cfg Config

	configPath := getConfigFilePath()

	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(configBytes, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) SetUser(user_name string) error {
	// Export a SetUser method on the Config struct that writes the config struct to the JSON file after setting the current_user_name field.'
	cfg.CurrentUser = user_name
	if err := write(*cfg); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func getConfigFilePath() string {
	homeDir, _ := os.UserHomeDir()
	configPath := filepath.Join(homeDir, ".gatorconfig.json")
	return configPath
}

func write(cfg Config) error {
	// Marshal the struct to JSON
	cfgBytes, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	// Write the JSON bytes to a file
	configPath := getConfigFilePath()

	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.Write(cfgBytes); err != nil {
		return err
	}

	return nil
}
