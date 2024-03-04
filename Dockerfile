FROM golang:1.22.0-alpine

WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod download

COPY cmd/ /app/cmd/
COPY internal/ /app/internal/
COPY docs/ /app/docs
COPY .env.example .env

RUN go build -o migration transaction-system/cmd/migration

RUN go build -o main transaction-system/cmd/transaction-system

COPY entrypoint.sh /app/entrypoint.sh

RUN chmod +x /app/entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]