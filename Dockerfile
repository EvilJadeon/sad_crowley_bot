FROM golang:1.24.1-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o crowley ./internal/app/bot && \
    chmod +x crowley

EXPOSE 8080

CMD ["./crowley"]
