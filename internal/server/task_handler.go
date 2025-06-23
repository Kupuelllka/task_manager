package server

import "net/http"

// TaskHandlers определяет интерфейс для обработчиков задач
type TaskHandlers interface {
	CreateTaskHandler(w http.ResponseWriter, r *http.Request)
	GetTaskHandler(w http.ResponseWriter, r *http.Request)
	DeleteTaskHandler(w http.ResponseWriter, r *http.Request)
}
