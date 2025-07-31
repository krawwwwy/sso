FROM golang:1.24-alpine AS builder
=======
FROM golang:1.21-alpine AS builder
>>>>>>> 381994eea2fc54ef6f274240221f0099c08e4c65

WORKDIR /app

# Копируем модули
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходники
COPY . .

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sso ./cmd/sso

<<<<<<< HEAD
FROM golang:1.24-alpine
=======
FROM golang:1.21-alpine
>>>>>>> 381994eea2fc54ef6f274240221f0099c08e4c65

RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app

# Копируем бинарник
COPY --from=builder /app/sso .

# Копируем исходники для миграций (нужен Go)
COPY . .

# Создаем директории
RUN mkdir -p /app/config /app/storage

EXPOSE 44044

CMD ["./sso"] 