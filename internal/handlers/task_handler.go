package handlers

import (
	"encoding/json"
	"net/http"
	task "taskmanager/internal/service/task_service"
	"taskmanager/pkg/logger"
)

// APIHandler реализует интерфейс TaskHandlers
type TaskHandler struct {
	service *task.TaskService
	logger  *logger.Log
}

// NewAPIHandler создает новый обработчик API
func NewTaskHandler(service *task.TaskService, logger *logger.Log) *TaskHandler {
	return &TaskHandler{
		service: service,
		logger:  logger,
	}
}

// CreateTaskHandler обрабатывает создание задачи
func (h *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.logger.Debugf("Invalid method for task creation: %s", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	task := h.service.CreateTask()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(task); err != nil {
		h.logger.Errorf("Failed to encode task response: %v", err)
	}
}

// GetTaskHandler возвращает информацию о задаче
func (h *TaskHandler) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.logger.Debugf("Invalid method for task retrieval: %s", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	taskID := r.URL.Query().Get("id")
	if taskID == "" {
		h.logger.Debug("Task ID not provided in request")
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	task, exists := h.service.GetTask(taskID)
	if !exists {
		h.logger.Debugf("Task not found: %s", taskID)
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		h.logger.Errorf("Failed to encode task response: %v", err)
	}
}

// DeleteTaskHandler удаляет задачу
func (h *TaskHandler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		h.logger.Debugf("Invalid method for task deletion: %s", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	taskID := r.URL.Query().Get("id")
	if taskID == "" {
		h.logger.Debug("Task ID not provided in request")
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	deleted := h.service.DeleteTask(taskID)
	if !deleted {
		h.logger.Debugf("Task not found for deletion: %s", taskID)
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
