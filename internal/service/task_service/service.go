package task

import (
	"fmt"
	"sync"
	"taskmanager/pkg/logger"
	"time"
)

// TaskService управляет задачами
type TaskService struct {
	mu     sync.RWMutex
	tasks  map[string]*Task
	logger *logger.Log
}

// NewTaskService создает новый сервис задач
func NewTaskService(logger *logger.Log) *TaskService {
	return &TaskService{
		tasks:  make(map[string]*Task),
		logger: logger,
	}
}

// CreateTask создает новую задачу
func (s *TaskService) CreateTask() *Task {
	taskID := fmt.Sprintf("task_%d", time.Now().UnixNano())

	task := &Task{
		ID:        taskID,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	s.mu.Lock()
	s.tasks[taskID] = task
	s.mu.Unlock()

	s.logger.Infof("Created new task with ID: %s", taskID)

	// Запускаем обработку задачи в фоне
	go s.processTask(task)

	return task
}

// processTask имитирует долгую обработку задачи
func (s *TaskService) processTask(task *Task) {
	s.mu.Lock()
	task.Status = "processing"
	task.StartedAt = time.Now()
	s.mu.Unlock()

	s.logger.Debugf("Started processing task %s", task.ID)

	// Имитация долгой I/O bound задачи (3-5 минут)
	time.Sleep(time.Duration(3+time.Now().UnixNano()%3) * time.Minute)

	s.mu.Lock()
	defer s.mu.Unlock()

	task.Status = "completed"
	task.FinishedAt = time.Now()
	task.Duration = task.FinishedAt.Sub(task.StartedAt)
	task.Result = fmt.Sprintf("Результат обработки задачи %s", task.ID)

	s.logger.Infof("Completed task %s, duration: %v", task.ID, task.Duration)
}

// GetTask возвращает задачу по ID
func (s *TaskService) GetTask(id string) (*Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[id]
	if !exists {
		s.logger.Debugf("Task not found: %s", id)
		return nil, false
	}

	s.logger.Debugf("Retrieved task %s, status: %s", id, task.Status)
	return task, true
}

// DeleteTask удаляет задачу
func (s *TaskService) DeleteTask(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		s.logger.Debugf("Attempt to delete non-existent task: %s", id)
		return false
	}

	delete(s.tasks, id)
	s.logger.Infof("Deleted task: %s", id)
	return true
}
