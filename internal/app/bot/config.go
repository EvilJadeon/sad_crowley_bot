package bot

import "os"

// Config ...
type Config struct {
	TelegramBotToken string `env:"TELEGRAM_BOT_API_TOKEN"`
	LogLevel         string `env:"LOG_LEVEL"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_API_TOKEN"),
		LogLevel:         os.Getenv("LOG_LEVEL"),
	}
}
