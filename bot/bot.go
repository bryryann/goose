package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bryryann/goose/config"
	"github.com/bwmarrin/discordgo"
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

func Run() {
	discord, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal(err)
	}

	discord.Open()
	defer discord.Close()

	fmt.Println(asciiTitle)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
