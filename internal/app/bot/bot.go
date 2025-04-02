package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

// Bot ...
type Bot struct {
	config *Config
	api    *tgbotapi.BotAPI
	logger *logrus.Logger
}

// New ...
func New(config *Config) *Bot {
	return &Bot{
		config: config,
		api:    nil,
		logger: logrus.New(),
	}
}

func (b *Bot) SetLogger() error {
	level, err := logrus.ParseLevel(b.config.LogLevel)

	if err != nil {
		return err
	}

	b.logger.SetLevel(level)

	return nil
}

// Start ...
func (b *Bot) Start() error {
	if err := b.SetLogger(); err != nil {
		return err
	}

	bot, err := tgbotapi.NewBotAPI(b.config.TelegramBotToken)

	if err != nil {
		b.logger.Errorf("Failed to create bot: %v", err)
	}

	b.api = bot
	bot.Debug = true

	b.logger.Infof("Authorized on account %s", bot.Self.UserName)

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
		b.logger.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I received your message: "+update.Message.Text)
		if _, err := b.api.Send(msg); err != nil {
			b.logger.Errorf("Error sending message: %v", err)
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
		b.logger.Errorf("Error sending start message: %v", err)
	}
}

func (b *Bot) handleHelpCommand(update tgbotapi.Update) {
	helpText := `Available commands:
/start - Start the bot
/help - Show this help message`

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, helpText)
	if _, err := b.api.Send(msg); err != nil {
		b.logger.Errorf("Error sending help message: %v", err)
	}
}

func (b *Bot) handleUnknownCommand(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Unknown command. Use /help to see available commands.")
	if _, err := b.api.Send(msg); err != nil {
		b.logger.Errorf("Error sending unknown command message: %v", err)
	}
}
