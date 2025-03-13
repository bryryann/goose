package main

import (
	"log"

	"github.com/bryryann/goose/bot"
	"github.com/bryryann/goose/config"
	"github.com/joho/godotenv"
)

func init() { godotenv.Load() }

func main() {
	err := config.SetUp()
	if err != nil {
		log.Fatal(err)
	}

	bot.Run()
}
