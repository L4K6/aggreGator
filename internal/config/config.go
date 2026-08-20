package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	fullPath, err := getFilePath()
	if err != nil {
		return Config{}, err
	}
	var config Config
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return Config{}, err
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
func (c *Config) SetUser(user string) error {
	c.CurrentUserName = user
	jsonData, err := json.Marshal(c)
	if err != nil {
		return err
	}
	fullPath, err := getFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(fullPath, jsonData, 0666)
	if err != nil {
		return err
	}
	return nil
}
func getFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(homePath, configFileName)
	return fullPath, nil
}
