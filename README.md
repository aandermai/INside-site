# INside
Wiki-страница проекта: https://se.cs.petrsu.ru/wiki/INside

# Копировать репозиторий к себе
```
git clone https://github.com/aandermai/INside-site.git
```

## Требования
- Установленный Docker

## Как установить Docker на Windows (включая WSL)

### Вариант 1: Docker Desktop 
1. Скачай и установи Docker Desktop: https://www.docker.com/products/docker-desktop/
2. Запусти установщик и следуй инструкциям.
3. Во время установки поставь галочку «Use WSL 2 based engine».
4. Перезагрузи компьютер (если попросит).
5. Проверь: ``` docker version && docker compose version ```

Готово — можно запускать docker compose up.

### Вариант 2: Только WSL 2 + Docker внутри Linux
1. Выполни: ``` sudo apt update && sudo apt upgrade -y && sudo apt install docker.io docker-compose-plugin -y ```
2. Добавь себя в группу docker: ``` sudo usermod -aG docker $USER && newgrp docker ```
3. Проверь: ``` docker version && docker compose version ```

Теперь в терминале WSL можно запускать проект.

## Как запустить проект (самый простой способ — через Docker)
1. Скачай проект (git clone или ZIP)
2. В папке проекта выполни: mkdir data (один раз)
3. Запусти одной командой: ``` docker compose up --build ```
4. Открой в браузере: http://localhost:9000

Готово! Все регистрации сохраняются в папке data навсегда.

### Остановка
```
docker compose down
```

## Запуск без Docker (для разработчиков)
1. Установи Go: https://go.dev/dl/
2. В папке проекта: ``` go run ./backend/main.go ```
3. Открой http://localhost:9000

База создастся автоматически в папке data.

## Структура проекта
- backend/ — сервер (main.go, handlers, db)
- frontend/ — HTML, CSS, JS
- data/ — файл базы данных (создаётся автоматически)

## Функционал
- Регистрация и вход
- Профиль (имя, фамилия, email)
- Выход из аккаунта
- Данные сохраняются между перезапусками
