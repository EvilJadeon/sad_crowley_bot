package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Bot ...
type Bot struct {
	config *Config
	api    *tgbotapi.BotAPI
}

// New ...
func New(config *Config) *Bot {
	return &Bot{
		config: config,
		api:    nil,
	}
}

// Start ...
func (b *Bot) Start() error {
	bot, err := tgbotapi.NewBotAPI(b.config.TelegramBotToken)
	if err != nil {
		log.Fatal(err)
	}

	b.api = bot
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		b.handleMessage(update)
		b.handleCommand(update)
	}

	return nil
}

func (b *Bot) handleMessage(update tgbotapi.Update) {
	if update.Message != nil {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I received your message: "+update.Message.Text)
		if _, err := b.api.Send(msg); err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}
}

func (b *Bot) handleCommand(update tgbotapi.Update) {
	if update.Message != nil && update.Message.IsCommand() {
		switch update.Message.Command() {
		case "start":
			b.handleStartCommand(update)
		case "help":
			b.handleHelpCommand(update)
		default:
			b.handleUnknownCommand(update)
		}
	}
}

func (b *Bot) handleStartCommand(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome! I'm your bot. Use /help to see available commands.")
	if _, err := b.api.Send(msg); err != nil {
		log.Printf("Error sending start message: %v", err)
	}
}

func (b *Bot) handleHelpCommand(update tgbotapi.Update) {
	helpText := `Available commands:
/start - Start the bot
/help - Show this help message`

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, helpText)
	if _, err := b.api.Send(msg); err != nil {
		log.Printf("Error sending help message: %v", err)
	}
}

func (b *Bot) handleUnknownCommand(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Unknown command. Use /help to see available commands.")
	if _, err := b.api.Send(msg); err != nil {
		log.Printf("Error sending unknown command message: %v", err)
	}
}
