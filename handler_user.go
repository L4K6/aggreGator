package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/L4K6/aggreGator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, c command) error {
	if len(c.args) == 0 {
		return errors.New("login name missing")
	}
	if _, err := s.queries.GetUser(context.Background(), c.args[0]); err != nil {
		fmt.Println("user doesnt exist: ", err)
		os.Exit(1)

	}
	s.config.CurrentUserName = c.args[0]
	s.config.SetUser(s.config.CurrentUserName)
	fmt.Println("user has been logged in")
	return nil
}

func handlerRegister(s *state, c command) error {
	if len(c.args) == 0 {
		return errors.New("register name missing")
	}
	newUUID := uuid.New()
	newUser := database.CreateUserParams{ID: newUUID, CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: c.args[0]}
	dbUser, err := s.queries.CreateUser(context.Background(), newUser)
	if err != nil {
		fmt.Println("error creating user:", err)
		os.Exit(1)
	}
	s.config.CurrentUserName = c.args[0]
	s.config.SetUser(s.config.CurrentUserName)
	fmt.Println("user has been created")
	log.Println(dbUser)
	return nil
}
