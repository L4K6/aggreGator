package main

import (
	"fmt"
	"log"

	"github.com/L4K6/aggreGator.git/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading a file: %v", err)
	}
	err = cfg.SetUser("Crane")
	if err != nil {
		log.Fatalf("Error setting a user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error reading a file: %v", err)
	}
	fmt.Print(cfg)
}
