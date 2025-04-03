package bot

import (
	"os"

	"github.com/EvilJadeon/sad_crowley_bot/internal/app/store"
)

// Config ...
type Config struct {
	TelegramBotToken string `env:"TELEGRAM_BOT_API_TOKEN"`
	LogLevel         string `env:"LOG_LEVEL"`
	Store            *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_API_TOKEN"),
		LogLevel:         os.Getenv("LOG_LEVEL"),
		Store:            store.NewConfig(),
	}
}
