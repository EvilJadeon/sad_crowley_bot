package main

import (
	"log"
	"os"

	"github.com/EvilJadeon/sad_crowley_bot/internal/app/bot"
)

func main() {
	config := bot.NewConfig()
	config.TelegramBotToken = os.Getenv("TELEGRAM_BOT_API_TOKEN")

	b := bot.New(config)
	if err := b.Start(); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}
}
