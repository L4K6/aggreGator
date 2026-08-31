package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/L4K6/aggreGator/internal/config"
	"github.com/L4K6/aggreGator/internal/database"
	_ "github.com/lib/pq"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	queries := database.New(db)

	stateInstance := state{config: &cfg, queries: queries}
	commands := commands{commands: make(map[string]func(s *state, c command) error)}

	systemArguments := os.Args
	if len(systemArguments) < 2 {
		fmt.Fprintln(os.Stderr, "incorrect amount of arguments supplied")
		os.Exit(1)
	}

	cmd := command{name: systemArguments[1], args: systemArguments[2:]}

	commands.register("login", func(stateInstance *state, cmd command) error { return handlerLogin(stateInstance, cmd) })
	commands.register("register", func(stateInstance *state, cmd command) error { return handlerRegister(stateInstance, cmd) })
	commands.register("reset", func(stateInstance *state, cmd command) error { return commands.reset(stateInstance, cmd) })
	commands.register("users", func(stateInstance *state, cmd command) error { return commands.users(stateInstance, cmd) })
	commands.register("agg", func(stateInstance *state, cmd command) error { return agg(stateInstance, cmd) })
	commands.register("addfeed", func(stateInstance *state, cmd command) error { return addfeed(stateInstance, cmd) })
	commands.register("feeds", func(stateInstance *state, cmd command) error { return feeds(stateInstance, cmd) })

	err = commands.run(&stateInstance, cmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
