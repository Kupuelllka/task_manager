package main

import (
	"fmt"
	"log"
	"taskmanager/config"
	"taskmanager/internal/handlers"
	"taskmanager/internal/server"
	task "taskmanager/internal/service/task_service"
	"taskmanager/pkg/logger"
)

func main() {
	cfg, err := config.GetConfig("./config.yml")
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}
	// Инициализация логгера
	logger, err := logger.NewLogger("all", "debug", "task_service.log")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	// Инициализация зависимостей
	service := task.NewTaskService(logger)
	handler := handlers.NewTaskHandler(service, logger)
	server := server.NewServer(handler, logger)

	logger.Infof("task service started on %s:%s", cfg.General.Host, cfg.General.Port)
	// Запуск сервера
	if err := server.Start(fmt.Sprintf("%s:%s", cfg.General.Host, cfg.General.Port)); err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
}
