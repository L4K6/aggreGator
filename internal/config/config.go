package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

type state struct {
	config *Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	commands map[string]func(s *state, c command) error
}

func (cmds *commands) run(s *state, c command) error {
	function, ok := cmds.commands[c.name]
	if !ok {
		return errors.New("command doesnt exist")
	}
	err := function(s, c)
	if err != nil {
		return err
	}
	return nil
}

func (cmds *commands) register(name string, f func(*state, command) error) {
	cmds.commands[name] = f
}

func handlerLogin(s *state, c command) error {
	if len(c.args) == 0 {
		return errors.New("login name missing")
	} else {
		s.config.CurrentUserName = c.args[0]
		fmt.Println("user has been set")
		return nil
	}
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
