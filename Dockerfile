# 🔥 Используем обычный образ Golang (не Alpine)
FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /crowley ./internal/app/bot

FROM debian:stable-slim

WORKDIR /app

COPY --from=builder /crowley /app/crowley

RUN chmod +x /app/crowley

EXPOSE 8080

CMD ["./crowley"]