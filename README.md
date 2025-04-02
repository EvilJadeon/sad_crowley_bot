# Sad Crowley Bot

A Telegram bot built with Go.

## Features

- Responds to messages
- Handles basic commands (/start, /help)
- Built with Go and the Telegram Bot API

## Prerequisites

- Go 1.24.1 or later
- Docker and Docker Compose
- Telegram Bot Token

## Setup

1. Clone the repository
2. Copy `.env.example` to `.env` and fill in your Telegram Bot Token
3. Run with Docker Compose:
   ```bash
   docker compose up -d
   ```

## Development

To run locally:

```bash
make run
```

## License

MIT
