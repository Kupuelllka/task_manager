package server

import (
	"net/http"
	"net/http/httptest"
	"taskmanager/internal/handlers"
	task "taskmanager/internal/service/task_service"
	"taskmanager/pkg/logger"
	"testing"
	"time"
)

func TestTaskService_CreateTask(t *testing.T) {
	logger, _ := logger.NewLogger("console", "debug", "")
	service := task.NewTaskService(logger)
	task := service.CreateTask()

	if task.ID == "" {
		t.Error("Task ID should not be empty")
	}
	if task.Status != "pending" {
		t.Errorf("Expected status 'pending', got '%s'", task.Status)
	}
}

func TestAPIHandler_CreateTaskHandler(t *testing.T) {
	logger, _ := logger.NewLogger("console", "debug", "")
	service := task.NewTaskService(logger)
	handler := handlers.NewTaskHandler(service, logger)

	req := httptest.NewRequest("POST", "/tasks/create", nil)
	w := httptest.NewRecorder()

	handler.CreateTaskHandler(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestAPIHandler_GetTaskHandler(t *testing.T) {
	logger, _ := logger.NewLogger("console", "debug", "")
	service := task.NewTaskService(logger)
	handler := handlers.NewTaskHandler(service, logger)

	// Сначала создаем задачу
	task := service.CreateTask()

	// Теперь запрашиваем ее статус
	req := httptest.NewRequest("GET", "/tasks/get?id="+task.ID, nil)
	w := httptest.NewRecorder()

	handler.GetTaskHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestAPIHandler_DeleteTaskHandler(t *testing.T) {
	logger, _ := logger.NewLogger("console", "debug", "")
	service := task.NewTaskService(logger)
	handler := handlers.NewTaskHandler(service, logger)

	// Сначала создаем задачу
	task := service.CreateTask()

	// Теперь удаляем ее
	req := httptest.NewRequest("DELETE", "/tasks/delete?id="+task.ID, nil)
	w := httptest.NewRecorder()

	handler.DeleteTaskHandler(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, got %d", http.StatusNoContent, w.Code)
	}
}

func TestTask_Lifecycle(t *testing.T) {
	logger, _ := logger.NewLogger("console", "debug", "")
	service := task.NewTaskService(logger)

	// Создаем задачу
	task := service.CreateTask()
	time.Sleep(100 * time.Millisecond) // Даем время на старт обработки

	// Проверяем что задача перешла в состояние processing
	storedTask, exists := service.GetTask(task.ID)
	if !exists {
		t.Fatal("Task should exist")
	}
	if storedTask.Status != "processing" {
		t.Errorf("Expected status 'processing', got '%s'", storedTask.Status)
	}

	// Удаляем задачу
	deleted := service.DeleteTask(task.ID)
	if !deleted {
		t.Error("Task should be deleted successfully")
	}
}
