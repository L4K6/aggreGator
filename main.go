package main

import (
	"fmt"
	"os"

	"github.com/L4K6/aggreGator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	stateInstance := state{config: &cfg}
	commands := commands{commands: make(map[string]func(s *state, c command) error)}

	systemArguments := os.Args
	if len(systemArguments) < 2 {
		fmt.Fprintln(os.Stderr, "incorrect amount of arguments supplied")
		os.Exit(1)
	}

	cmd := command{name: systemArguments[1], args: systemArguments[2:]}

	commands.register("login", func(stateInstance *state, cmd command) error { return handlerLogin(stateInstance, cmd) })

	err = commands.run(&stateInstance, cmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
