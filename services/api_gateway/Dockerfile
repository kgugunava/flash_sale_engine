# Stage 1: Build
FROM golang:1.23-alpine AS builder

# Установка зависимостей для сборки
RUN apk add --no-cache git gcc musl-dev

# Устанавливаем рабочую директорию
WORKDIR /build

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код всего проекта (нужно для импортов из /shared и /pkg)
COPY . /build/

# Собираем приложение
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build \
    -ldflags "-w -s" \
    -a -installsuffix cgo \
    -o api_gateway ./services/api_gateway/cmd/main

# Stage 2: Runtime
FROM alpine:latest

# Установка runtime зависимостей
RUN apk --no-cache add ca-certificates tzdata

# Создаем пользователя для безопасности
RUN addgroup -g 1000 apigateway && \
    adduser -D -u 1000 -G apigateway apigateway

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем бинарник из builder stage
COPY --from=builder /build/api_gateway .

# Переходим на непривилегированного пользователя
USER apigateway

# Exposing port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=40s --retries=3 \
    CMD wget --quiet --tries=1 --spider http://localhost:8080/health || exit 1

# Запуск приложения
CMD ["./api_gateway"]
