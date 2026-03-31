# ---------- STAGE 1: build ----------
FROM golang:1.24 AS builder

WORKDIR /app

# кешируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# копируем весь проект
COPY . .

# собираем бинарник
RUN CGO_ENABLED=1 GOOS=linux go build -o app cmd/main.go

# ---------- STAGE 2: runtime ----------
FROM debian:bookworm-slim

WORKDIR /app

# копируем бинарник
COPY --from=builder /app/app .

# копируем конфиги
COPY config ./config

# (опционально) если используешь sqlite файл
COPY storage ./storage

# Для миграций
COPY internal ./internal

# порт (для понимания, не обязателен)
EXPOSE 8082

# команда запуска
CMD ["./app"]