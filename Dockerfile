FROM golang:1.25.5

# Рабочая директория — корень проекта (как на твоей машине)
WORKDIR /app

# Копируем всё
COPY . .

# Переходим в папку с main.go для сборки
WORKDIR /app/backend

# Скачиваем зависимости и собираем
RUN go mod download && go build -o ../main .

# Возвращаемся в корень и запускаем
WORKDIR /app

EXPOSE 9000

CMD ["./main"]