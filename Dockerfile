# Build Stage
FROM golang:1.24 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /app/auth-service ./cmd/dco

# Execution Stage
FROM alpine:3.20

WORKDIR /app

COPY --from=build /app/auth-service ./auth-service

EXPOSE 8080

CMD ["./auth-service"]
