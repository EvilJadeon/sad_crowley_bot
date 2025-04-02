package main

import (
	"log"

	"github.com/EvilJadeon/sad_crowley_bot/internal/app/bot"
)

func main() {
	config := bot.NewConfig()

	b := bot.New(config)

	if err := b.Start(); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}
}
