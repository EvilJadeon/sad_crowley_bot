package bot

// Config ...
type Config struct {
	TelegramBotToken string `env:"TELEGRAM_BOT_API_TOKEN"`
}

func NewConfig() *Config {
	return &Config{}
}
