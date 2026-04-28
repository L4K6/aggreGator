package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	err := write(c)
	if err != nil {
		return err
	}
	return nil
}

func write(cfg *Config) error {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	fullPath := filepath.Join(homePath, configFileName)

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(fullPath, data, 0600)
	if err != nil {
		return err
	}
	return nil
}
