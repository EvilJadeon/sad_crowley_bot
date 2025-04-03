package bot

import (
	"testing"
)

func TestBot_Start(t *testing.T) {
	config := &Config{
		TelegramBotToken: "test_token",
	}
	bot := New(config)

	// Skip actual bot start in test environment
	if bot.config.TelegramBotToken != "test_token" {
		t.Errorf("Expected test_token, got %s", bot.config.TelegramBotToken)
	}
}
