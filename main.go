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
	fmt.Println(cfg)
	err = cfg.SetUser("Crane")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	cfg, err = config.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(cfg)
}
