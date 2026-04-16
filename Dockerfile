# ---------- STAGE 1: build ----------
FROM golang:1.24 AS builder

WORKDIR /app

# кешируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# копируем весь проект
COPY . .

# собираем бинарник
RUN CGO_ENABLED=1 GOOS=linux go build -o app ./cmd/

# ---------- STAGE 2: runtime ----------
FROM debian:bookworm-slim

WORKDIR /app

# копируем бинарник
COPY --from=builder /app/app .

# копируем конфиги
COPY config ./config

# Миграции
COPY internal ./internal

# Создаём директорию для SQLite базы данных
RUN mkdir -p ./storage

EXPOSE 8082

# команда запуска
CMD ["./app"]