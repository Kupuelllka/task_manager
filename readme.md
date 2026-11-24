
# Task Service API

REST API для управления долгоживущими I/O-bound задачами с хранением состояния в памяти.

---

## 📌 Возможности

- ✅ Создание асинхронных задач (время выполнения 3–5 минут)
- ✅ Просмотр статуса задач
- ✅ Удаление задач
- ✅ Подробное логирование операций
- ✅ In-memory хранение состояния

---

## 🚀 Быстрый старт

**Запуск сервера (порт `8080`):**

```bash
# Через go run
go run main.go

# Или через сборку бинарника
go build -o task-service && ./task-service
```

**Логи доступны:**

- в консоли;
- в файле `task_service.log`.

---

## 📡 API Endpoints

### 1. Создание задачи

| Параметр | Значение |
|----------|----------|
| **Метод** | `POST` |
| **Endpoint** | `/tasks/create` |

**Запрос:**

```http
POST /tasks/create HTTP/1.1
Host: localhost:8080
```

**Ответ:**

```http
HTTP/1.1 201 Created
Content-Type: application/json
```

```json
{
  "id": "task_123456789",
  "status": "pending",
  "created_at": "2023-06-20T14:30:00Z"
}
```

---

### 2. Получение статуса задачи

| Параметр | Значение |
|----------|----------|
| **Метод** | `GET` |
| **Endpoint** | `/tasks/get?id=<task_id>` |

**Запрос:**

```http
GET /tasks/get?id=task_123456789 HTTP/1.1
Host: localhost:8080
```

**Ответ (задача в процессе):**

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "id": "task_123456789",
  "status": "processing",
  "created_at": "2023-06-20T14:30:00Z",
  "started_at": "2023-06-20T14:30:01Z",
  "duration": 150000000000
}
```

**Ответ (задача завершена):**

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

---

### 3. Удаление задачи

| Параметр | Значение |
|----------|----------|
| **Метод** | `DELETE` |
| **Endpoint** | `/tasks/delete?id=<task_id>` |

**Запрос:**

```http
DELETE /tasks/delete?id=task_123456789 HTTP/1.1
Host: localhost:8080
```

**Ответ:**

```http
HTTP/1.1 204 No Content
```

---

## 🛠 Примеры использования

### cURL

```bash
# Создание задачи
curl -X POST http://localhost:8080/tasks/create

# Проверка статуса
curl "http://localhost:8080/tasks/get?id=task_123456789"

# Удаление задачи
curl -X DELETE "http://localhost:8080/tasks/delete?id=task_123456789"
```

### HTTPie

```bash
# Создание задачи
http POST :8080/tasks/create

# Проверка статуса
http GET :8080/tasks/get id==task_123456789

# Удаление задачи
http DELETE :8080/tasks/delete id==task_123456789
```

---

## 🧪 Тесты

**Запуск тестов:**

```bash
go test -v ./...
```

**Тесты покрывают:**

- создание задач;
- проверку статусов;
- удаление задач;
- жизненный цикл задачи;
- обработку ошибок.

---

## 📊 Логирование

**Уровни логирования** (настраиваются при запуске):

| Уровень | Описание |
|---------|----------|
| `error` | Только критические ошибки |
| `info` | Основные операции (`+error`) |
| `debug` | Детальная отладка (`+info + error`) |
| `debug-raw` | Максимальная детализация |

**Пример лога (`task_service.log`):**

```log
2023/06/20 14:30:00 [INFO]  Created task: task_123456789
2023/06/20 14:30:01 [DEBUG] Processing started: task_123456789
2023/06/20 14:33:45 [INFO]  Task completed: task_123456789 (duration: 3m45s)
```

---

## 🔮 Roadmap

- [ ] Graceful shutdown
- [ ] Prometheus-метрики
- [ ] JWT-аутентификация
- [ ] Поддержка Redis для хранения состояния
- [ ] WebSocket для push-уведомлений
