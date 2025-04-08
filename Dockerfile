# 🔥 Используем обычный образ Golang (не Alpine)
FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /crowley ./cmd/bot/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /crowley /app/crowley
COPY migrations /app/migrations

RUN chmod +x /app/crowley

# Install required packages including migrate
RUN apk add --no-cache ca-certificates curl tar

# Install migrate tool
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz \
    && mv migrate /usr/local/bin/migrate \
    && chmod +x /usr/local/bin/migrate

EXPOSE 8080

CMD ["./crowley"]