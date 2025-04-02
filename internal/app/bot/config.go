package bot

import "os"

// Config ...
type Config struct {
	TelegramBotToken string `env:"TELEGRAM_BOT_API_TOKEN"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_API_TOKEN"),
	}
}
