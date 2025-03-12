package main

import (
	"fmt"
	"log"

	"github.com/bryryann/goose/config"
	"github.com/joho/godotenv"
)

var (
	asciiTitle = `
  ________
 /  _____/  ____   ____  ______ ____
/   \  ___ /  _ \ /  _ \/  ___// __ \
\    \_\  (  <_> |  <_> )___ \\  ___/
 \________/\____/ \____/______\\____\
	`
)

func init() { godotenv.Load() }

func main() {
	err := config.SetUp()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(asciiTitle)
}
