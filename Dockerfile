FROM golang:1.21-alpine AS builder

WORKDIR /app

# Копируем go mod файлы
COPY go.mod ./
RUN go mod download

# Копируем исходный код
COPY *.go ./

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Финальный образ
FROM alpine:latest

WORKDIR /root/

# Копируем бинарник из builder
COPY --from=builder /app/app .

# Создаем непривилегированного пользователя
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Запускаем приложение
CMD ["./app"]

