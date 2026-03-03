#FROM golang:1.24-alpine AS builder
#
#WORKDIR /app
#
#COPY go.mod go.sum ./
#RUN go mod download
#
#COPY . .
#
#RUN go build -o auth-service ./cmd/dco
#
#FROM alpine:latest
#
#WORKDIR /app
#
#COPY --from=builder /app/auth-service .
#
#CMD ["./auth-service"]

# Этап сборки
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Копируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код
COPY . .

# Собираем приложение из папки cmd/dco
RUN go build -o auth-service ./cmd/dco/main.go

# Этап запуска
FROM alpine:latest
WORKDIR /root/

# Копируем скомпилированный файл и конфиг
COPY --from=builder /app/auth-service .
COPY --from=builder /app/config.yaml .

# Если вашему приложению нужны миграции или логи, убедитесь, что пути совпадают
EXPOSE 8080

CMD ["./auth-service"]