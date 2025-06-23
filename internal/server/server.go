package server

import (
	"net/http"
	"taskmanager/pkg/logger"
)

// Server представляет HTTP сервер
type Server struct {
	handlers TaskHandlers
	logger   *logger.Log
}

// NewServer создает новый экземпляр сервера
func NewServer(handlers TaskHandlers, logger *logger.Log) *Server {
	return &Server{
		handlers: handlers,
		logger:   logger,
	}
}

// Start запускает HTTP сервер
func (s *Server) Start(addr string) error {
	http.HandleFunc("/tasks/create", s.handlers.CreateTaskHandler)
	http.HandleFunc("/tasks/get", s.handlers.GetTaskHandler)
	http.HandleFunc("/tasks/delete", s.handlers.DeleteTaskHandler)

	s.logger.Infof("Server starting on %s...", addr)
	return http.ListenAndServe(addr, nil)
}
