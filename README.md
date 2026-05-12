# Backend (Gin + Swagger)

## Подготовка окружения

Перед запуском необходимо создать файл .env с переменными окружения.

Скопируйте пример конфигурации командой

    cp .env.example .env

При необходимости отредактируйте значения.


## Запуск Gin сервера и инфраструктуры

    docker compose up --build

После запуска будут подняты:
- Backend (Gin сервер)
- PostgreSQL
- MinIO

Backend будет доступен на http://localhost:8080


## Локальный запуск Gin сервера

    go run cmd/server/main.go

Сервер будет слушать на http://localhost:8080


## Миграции базы данных

Команда выполняется один раз при первом запуске проекта или после сброса данных PostgreSQL.

Установка migrate:

    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

Применение миграций (строка подключения зависит от переменных окружения):

    migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/skin_service?sslmode=disable" up


## Swagger UI

Swagger доступен по адресу http://localhost:8080/swagger/index.html

Генерация документации:

    swag init -g cmd/server/main.go


## Полезные команды

Пересборка Docker:

    docker compose down
    docker compose up --build

Запуск Docker:

    docker compose up

Полный сброс данных (PostgreSQL + MinIO):

    docker compose down -v
