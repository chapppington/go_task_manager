# Техническое задание: Task Management System

## 1. Общее описание

Система управления задачами на основе микросервисной архитектуры с использованием gRPC, dependency injection (Dig) и Clean Architecture.

## 2. Архитектура

### 2.1 Микросервисы

#### 2.1.1 User Service
- **Назначение**: Управление пользователями системы
- **Порт gRPC**: 50051
- **Порт HTTP (для health checks)**: 8001

#### 2.1.2 Task Service
- **Назначение**: Управление задачами пользователей
- **Порт gRPC**: 50052
- **Порт HTTP (для health checks)**: 8002
- **Зависимости**: User Service (для валидации существования пользователя)

### 2.2 Структура проекта

```
go_crud/
├── cmd/
│   ├── user-service/
│   │   └── main.go
│   └── task-service/
│       └── main.go
├── internal/
│   ├── user/
│   │   ├── domain/
│   │   │   ├── user.go          # Entity
│   │   │   └── repository.go    # Repository interface
│   │   ├── application/
│   │   │   └── service.go        # Business logic
│   │   ├── infrastructure/
│   │   │   ├── repository/
│   │   │   │   └── postgres.go  # PostgreSQL implementation
│   │   │   └── database.go
│   │   └── presentation/
│   │       └── grpc/
│   │           ├── handler.go   # gRPC handler
│   │           └── server.go
│   ├── task/
│   │   ├── domain/
│   │   │   ├── task.go
│   │   │   └── repository.go
│   │   ├── application/
│   │   │   └── service.go
│   │   ├── infrastructure/
│   │   │   ├── repository/
│   │   │   │   └── postgres.go
│   │   │   └── database.go
│   │   └── presentation/
│   │       └── grpc/
│   │           ├── handler.go
│   │           └── server.go
│   └── shared/
│       ├── grpc/
│       │   └── user_client.go   # gRPC client для User Service
│       └── errors/
│           └── errors.go
├── proto/
│   ├── user/
│   │   └── user.proto
│   └── task/
│       └── task.proto
├── config/
│   └── config.go
├── docker-compose.yml
├── Dockerfile
├── Makefile
└── go.mod
```

## 3. Модели данных

### 3.1 User (Пользователь)

```go
type User struct {
    ID        int64     `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

**Валидация:**
- Name: обязательное, 2-100 символов
- Email: обязательное, валидный email формат, уникальный

### 3.2 Task (Задача)

```go
type Task struct {
    ID          int64     `json:"id"`
    UserID      int64     `json:"user_id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Status      string    `json:"status"` // "todo", "in_progress", "done"
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

**Валидация:**
- UserID: обязательное, должен существовать в User Service
- Title: обязательное, 1-200 символов
- Description: опциональное, максимум 1000 символов
- Status: обязательное, один из: "todo", "in_progress", "done"

## 4. gRPC API

### 4.1 User Service API

#### 4.1.1 CreateUser
```protobuf
rpc CreateUser(CreateUserRequest) returns (UserResponse);
```

**Request:**
- name (string, required)
- email (string, required)

**Response:**
- id, name, email, created_at, updated_at

**Ошибки:**
- INVALID_ARGUMENT: невалидные данные
- ALREADY_EXISTS: email уже существует

#### 4.1.2 GetUser
```protobuf
rpc GetUser(GetUserRequest) returns (UserResponse);
```

**Request:**
- id (int64, required)

**Ошибки:**
- NOT_FOUND: пользователь не найден

#### 4.1.3 ListUsers
```protobuf
rpc ListUsers(ListUsersRequest) returns (ListUsersResponse);
```

**Request:**
- page (int32, optional, default: 1)
- page_size (int32, optional, default: 10)

**Response:**
- users (repeated UserResponse)
- total (int64)

#### 4.1.4 UpdateUser
```protobuf
rpc UpdateUser(UpdateUserRequest) returns (UserResponse);
```

**Request:**
- id (int64, required)
- name (string, optional)
- email (string, optional)

**Ошибки:**
- NOT_FOUND: пользователь не найден
- INVALID_ARGUMENT: невалидные данные
- ALREADY_EXISTS: email уже занят другим пользователем

#### 4.1.5 DeleteUser
```protobuf
rpc DeleteUser(DeleteUserRequest) returns (DeleteUserResponse);
```

**Request:**
- id (int64, required)

**Response:**
- success (bool)

**Ошибки:**
- NOT_FOUND: пользователь не найден

### 4.2 Task Service API

#### 4.2.1 CreateTask
```protobuf
rpc CreateTask(CreateTaskRequest) returns (TaskResponse);
```

**Request:**
- user_id (int64, required)
- title (string, required)
- description (string, optional)
- status (string, optional, default: "todo")

**Response:**
- id, user_id, title, description, status, created_at, updated_at

**Ошибки:**
- INVALID_ARGUMENT: невалидные данные
- NOT_FOUND: пользователь не найден (проверка через User Service)

#### 4.2.2 GetTask
```protobuf
rpc GetTask(GetTaskRequest) returns (TaskResponse);
```

**Request:**
- id (int64, required)

**Ошибки:**
- NOT_FOUND: задача не найдена

#### 4.2.3 ListTasks
```protobuf
rpc ListTasks(ListTasksRequest) returns (ListTasksResponse);
```

**Request:**
- user_id (int64, optional) - фильтр по пользователю
- status (string, optional) - фильтр по статусу
- page (int32, optional, default: 1)
- page_size (int32, optional, default: 10)

**Response:**
- tasks (repeated TaskResponse)
- total (int64)

#### 4.2.4 UpdateTask
```protobuf
rpc UpdateTask(UpdateTaskRequest) returns (TaskResponse);
```

**Request:**
- id (int64, required)
- title (string, optional)
- description (string, optional)
- status (string, optional)

**Ошибки:**
- NOT_FOUND: задача не найдена
- INVALID_ARGUMENT: невалидные данные

#### 4.2.5 DeleteTask
```protobuf
rpc DeleteTask(DeleteTaskRequest) returns (DeleteTaskResponse);
```

**Request:**
- id (int64, required)

**Response:**
- success (bool)

**Ошибки:**
- NOT_FOUND: задача не найдена

## 5. Технические требования

### 5.1 Dependency Injection (Dig)

Использовать `go.uber.org/dig` для управления зависимостями:

- Репозитории
- Сервисы
- gRPC handlers
- Database connections
- gRPC clients (для межсервисного взаимодействия)

### 5.2 База данных

- PostgreSQL для каждого сервиса (или общая БД с разными схемами)
- Миграции с помощью `golang-migrate` или `goose`
- Connection pooling

### 5.3 Обработка ошибок

- Стандартизированные gRPC статусы
- Логирование ошибок
- Валидация входных данных

### 5.4 Логирование

- Структурированное логирование (zap или logrus)
- Уровни: DEBUG, INFO, WARN, ERROR

### 5.5 Health Checks

- HTTP endpoint `/health` для каждого сервиса
- Проверка подключения к БД

## 6. Развертывание

### 6.1 Docker Compose

- User Service
- Task Service
- PostgreSQL (один или два инстанса)
- Возможность запуска через `docker-compose up`

### 6.2 Переменные окружения

**User Service:**
- `USER_SERVICE_PORT` (gRPC, default: 50051)
- `USER_HTTP_PORT` (HTTP, default: 8001)
- `POSTGRES_HOST`
- `POSTGRES_PORT`
- `POSTGRES_DB`
- `POSTGRES_USER`
- `POSTGRES_PASSWORD`
- `USER_SERVICE_HOST` (для Task Service)

**Task Service:**
- `TASK_SERVICE_PORT` (gRPC, default: 50052)
- `TASK_HTTP_PORT` (HTTP, default: 8002)
- `POSTGRES_HOST`
- `POSTGRES_PORT`
- `POSTGRES_DB`
- `POSTGRES_USER`
- `POSTGRES_PASSWORD`
- `USER_SERVICE_HOST` (адрес User Service)

## 7. Примеры использования

### 7.1 Создание пользователя
```bash
grpcurl -plaintext -d '{"name": "John Doe", "email": "john@example.com"}' \
  localhost:50051 user.UserService/CreateUser
```

### 7.2 Создание задачи
```bash
grpcurl -plaintext -d '{"user_id": 1, "title": "Learn gRPC", "description": "Study gRPC basics"}' \
  localhost:50052 task.TaskService/CreateTask
```

## 8. Этапы разработки

1. ✅ Настройка структуры проекта
2. ⬜ Настройка Dig контейнера
3. ⬜ Создание proto файлов и генерация кода
4. ⬜ Реализация User Service (domain, repository, service, handler)
5. ⬜ Реализация Task Service (domain, repository, service, handler)
6. ⬜ Настройка gRPC клиента для межсервисного взаимодействия
7. ⬜ Настройка базы данных и миграций
8. ⬜ Docker Compose конфигурация
9. ⬜ Тестирование
10. ⬜ Документация

## 9. Дополнительные возможности (опционально)

- Метрики (Prometheus)
- Трейсинг (Jaeger)
- API Gateway
- Аутентификация и авторизация
- WebSocket для real-time обновлений
- Кэширование (Redis)
