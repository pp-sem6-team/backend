# Backend (Gin + Swagger)

## Быстрый старт

### Запуск сервера

    go run cmd/server/main.go
Сервер будет слушать на :8080.

### Swagger UI
Доступен по адресу http://localhost:8080/swagger/index.html. \
Генерация документации после изменений в комментариях:

    swag init -g cmd/server/main.go
