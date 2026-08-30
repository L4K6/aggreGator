package main

import (
	"context"
	"errors"
	"fmt"

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

func (cmds *commands) reset(s *state, _ command) error {
	err := s.queries.DeleteAllUsers(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (cmds *commands) users(s *state, _ command) error {
	users, err := s.queries.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if s.config.CurrentUserName == user.Name {
			fmt.Println("*", user.Name, "(current)")
		} else {
			fmt.Println("*", user.Name)
		}
	}
	return nil
}

func agg(_ *state, _ command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", feed)
	return nil
}
func addfeed(s *state, c command) error {
	if len(c.args) != 2 {
		return fmt.Errorf("Missing argument feed")
	}
	feedName := c.args[0]
	url := c.args[1]
	userName := s.config.CurrentUserName

	user, err := s.queries.GetUser(context.Background(), userName)
	if err != nil {
		return err
	}
	dbArgs := database.CreateFeedParams{
		Name:   feedName,
		Url:    url,
		UserID: user.ID,
	}
	feed, err := s.queries.CreateFeed(context.Background(), dbArgs)
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}
