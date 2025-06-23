# Task Service API

REST API для управления долгоиграющими I/O bound задачами с хранением в памяти

## 📌 Основные возможности

- Создание асинхронных задач (3-5 мин выполнение)
- Просмотр статуса задач
- Удаление задач
- Подробное логирование операций
- In-memory хранение состояния

## 🚀 Быстрый старт

```bash
# Запуск сервера (порт 8080)
go run main.go

# Или с сборкой
go build -o task-service && ./task-service
Логи доступны:
```
В консоли

В файле task_service.log

📡 API Endpoints
1. Создание задачи
Endpoint:
POST /tasks/create

Request:
POST /tasks/create HTTP/1.1
Host: localhost:8080

Response:
HTTP/1.1 201 Created
Content-Type: application/json
```json
{
  "id": "task_123456789",
  "status": "pending",
  "created_at": "2023-06-20T14:30:00Z"
}
```
2. Получение статуса задачи
Endpoint:
GET /tasks/get?id=<task_id>

Request:
GET /tasks/get?id=task_123456789 HTTP/1.1
Host: localhost:8080

Response (в процессе):
HTTP/1.1 200 OK
Content-Type: application/json
```json
{
  "id": "task_123456789",
  "status": "processing",
  "created_at": "2023-06-20T14:30:00Z",
  "started_at": "2023-06-20T14:30:01Z",
  "duration": 150000000000
}
```
Response (завершено):

```json
{
  "id": "task_123456789",
  "status": "completed",
  "created_at": "2023-06-20T14:30:00Z",
  "started_at": "2023-06-20T14:30:01Z",
  "finished_at": "2023-06-20T14:33:45Z",
  "duration": 225000000000,
  "result": "Результат обработки задачи task_123456789"
}
```
3. Удаление задачи
Endpoint:
DELETE /tasks/delete?id=<task_id>

Request:
DELETE /tasks/delete?id=task_123456789 HTTP/1.1
Host: localhost:8080
Response:
HTTP/1.1 204 No Content

# 🛠 Примеры использования
cURL
bash
## Создание задачи
curl -X POST http://localhost:8080/tasks/create

## Проверка статуса
curl "http://localhost:8080/tasks/get?id=task_123456789"

## Удаление задачи
curl -X DELETE "http://localhost:8080/tasks/delete?id=task_123456789"

## Создание задачи
http POST :8080/tasks/create

## Проверка статуса
http GET :8080/tasks/get id==task_123456789

## Удаление задачи
http DELETE :8080/tasks/delete id==task_123456789

# Тесты
Запуск тестов:

```bash
go test -v ./...
```
## Тесты покрывают:
- Создание задач
- Проверку статусов
- Удаление задач
- Жизненный цикл задачи
- Обработку ошибок

# 📊 Логирование
## Уровни логирования (настраиваются при запуске):

- error    - Только критические ошибки
- info     - Основные операции (+error)
- debug    - Детальная отладка (+info+error)
- debug-raw- Максимальная детализация

Пример лога:

``` service.log
2023/06/20 14:30:00 [INFO] Created task: task_123456789
2023/06/20 14:30:01 [DEBUG] Processing started: task_123456789
2023/06/20 14:33:45 [INFO] Task completed: task_123456789 (duration: 3m45s)
```

# 🔮 Roadmap
- Добавить graceful shutdown
- Реализовать Prometheus метрики
- Добавить JWT аутентификацию
- Поддержка Redis для хранения состояния
- WebSocket для push-уведомлений

