FROM golang:1.24
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main ./cmd/app/main.go
# Устанавливаем dockerize
RUN wget https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz && \
    tar -xvzf dockerize-linux-amd64-v0.6.1.tar.gz && \
    mv dockerize /usr/local/bin/

# Переменные окружения
COPY .env .
# Ждем готовности сервисов и запускаем приложение
CMD ["dockerize", "-wait", "tcp://db:5433", "-timeout", "60s", "./main"]