package myApp

import (
	"log"
	"worm/internal/config"
)

// Определим структуру приложения, которая будет содержать зависимости для
// обработчиков HTTP, вспомогательных функций и middleware
type Application struct {
	Config *config.Config
	Logs   *log.Logger
}

func New(cfg *config.Config, logger *log.Logger) *Application {
	return &Application{
		Config: cfg,
		Logs:   logger,
	}
}
