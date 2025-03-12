package config

import (
	"fmt"
	"log"
	"os"
)

var (
	Token     string
	BotPrefix string
)

func SetUp() error {
	fmt.Println("setting up...")

	tokenEnv, exists := os.LookupEnv("TOKEN")
	if !exists {
		log.Fatal("Could not find TOKEN environment variable.")
	}
	Token = tokenEnv

	prefixEnv, exists := os.LookupEnv("PREFIX")
	if !exists {
		log.Fatal("Could not find PREFIX environment variable.")
	}
	BotPrefix = prefixEnv

	return nil
}
