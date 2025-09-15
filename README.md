# Checkers Game Server

Проект представляет собой сервер для игры в шашки, реализованный на Go с использованием фреймворка Gin и Swagger для документирования API.

## Структура проекта

```
checkers/
├── cmd/server/
│   └── main.go          # Основная точка входа приложения
├── internal/
│   ├── server/
|   |   ├── models.go    # Модели HTTP-запросов
│   │   └── handlers.go  # Обработчики HTTP-запросов
│   └── swagger/
│       └── swagger.go   # Конфигурация Swagger документации
├── pkg/
│   └── logger/          # Пакет для логирования
└── docs/                # Автогенерируемая документация Swagger
```

## Функциональность

- RESTful API для управления игрой в шашки
- Автоматическая генерация документации API через Swagger
- Логирование событий сервера
- Гибкая конфигурация запуска (с Swagger или без)

## API Endpoints

### POST /games
Создает новую игру в шашки.

**Параметры запроса:**
```json
{
  "player1": "string",
  "player2": "string"
}
```

**Ответ:**
```json
{
  "id": "string",
  "status": "string",
  "board": [[int]],
  "current_turn": "string",
  "winner": "string|null"
}
```

## Зависимости

- Go 1.24.6
- Gin Web Framework
- Swaggo (для генерации Swagger документации)

## Установка и запуск

1. Клонируйте репозиторий:
```bash
git clone <repository-url>
cd checkers
```

2. Установите зависимости:
```bash
go mod download
```

3. Установите Swag CLI для генерации документации:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

4. Запустите сервер:

**Без Swagger документации:**
```bash
go run cmd/server/main.go
```

**С Swagger документацией:**
```bash
ENABLE_SWAGGER=true go run -tags swagger cmd/server/main.go
```

После запуска с включенным Swagger, документация API будет доступна по адресу:
```
http://localhost:8080/swagger/index.html
```

## Генерация документации

Для ручной генерации Swagger документации выполните:
```bash
swag init -d cmd/server,internal/server
```

## Использование Makefile

Проект включает Makefile для упрощения запуска:

```bash
# Запуск без Swagger
make run

# Запуск с Swagger
make run-swagger
```

## Конфигурация

- Порт сервера: 8080
- Логирование настроено через пакет logger
- Swagger документация включается через переменную окружения `ENABLE_SWAGGER=true`
