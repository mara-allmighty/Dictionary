# Сборка
FROM golang:1.22 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN cd cmd && go build -o ../main .

# Запуск
FROM ubuntu:22.04

WORKDIR /root/

# Установим CA-сертификаты и netcat (для проверки подключения)
RUN apt-get update && \
    apt-get install -y ca-certificates netcat && \
    rm -rf /var/lib/apt/lists/*

# Копируем бинарник И конфиг из этапа builder
COPY --from=builder /app/main .
COPY --from=builder /app/config.yaml .

EXPOSE 8000
CMD ["./main"]
