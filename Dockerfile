FROM golang:1.24-alpine AS builder

WORKDIR /app

# Копируем модули
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходники
COPY . .

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sso ./cmd/sso

FROM golang:1.24-alpine

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