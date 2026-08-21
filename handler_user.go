package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, c command) error {
	if len(c.args) == 0 {
		return errors.New("login name missing")
	} else {
		s.config.CurrentUserName = c.args[0]
		s.config.SetUser(s.config.CurrentUserName)
		fmt.Println("user has been set")
		return nil
	}
}
