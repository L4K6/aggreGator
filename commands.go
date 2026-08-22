package main

import (
	"errors"

	"github.com/L4K6/aggreGator/internal/config"
	"github.com/L4K6/aggreGator/internal/database"
)

type state struct {
	config  *config.Config
	queries *database.Queries
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
