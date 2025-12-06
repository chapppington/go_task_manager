# Go Task Manager CRUD API

REST API для управления задачами и пользователями на Go.

## Технологии

- Go 1.25+
- Chi Router
- GORM + PostgreSQL
- Uber Dig (DI)
- Docker Compose

## Быстрый старт

1. Настройте `.env` файл:
```bash
cp .env.example .env
```

2. Запустите приложение:

```bash
# Запустить всё (БД + приложение)
make all

# Логи
make app-logs
```

## API

Базовый URL: `http://localhost:8000/api/v1`

### Пользователи
- `GET /users` - список пользователей
- `GET /users/{id}` - получить пользователя
- `GET /users/email/{email}` - найти по email
- `POST /users` - создать пользователя
- `PUT /users/{id}` - обновить пользователя
- `DELETE /users/{id}` - удалить пользователя

### Задачи
- `GET /tasks` - список задач
- `GET /tasks/{id}` - получить задачу
- `POST /tasks` - создать задачу
- `PUT /tasks/{id}` - обновить задачу
- `DELETE /tasks/{id}` - удалить задачу

### Health Check
- `GET /health` - проверка работоспособности

## Тестирование

```bash
make test
```

## Структура проекта

```
├── cmd/              # Точка входа
├── config/           # Конфигурация
├── internal/
│   ├── application/  # Use cases
│   ├── domain/       # Доменная модель
│   ├── infrastructure/ # Репозитории, БД
│   └── presentation/ # API handlers
└── tests/            # Тесты
```

## Make команды

- `make all` - запуск через Docker Compose
- `make all-down` - остановка всех контейнеров
- `make app-logs` - логи приложения
- `make postgres` - подключение к PostgreSQL
- `make test` - запуск тестов
