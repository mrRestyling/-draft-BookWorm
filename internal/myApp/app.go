package myApp

import (
	"log"
)

type Configurator interface {
	ParseFlags()
}

// Определим структуру приложения, которая будет содержать зависимости для
// обработчиков HTTP, вспомогательных функций и middleware
type Application struct {
	Config Configurator
	Logs   *log.Logger
}

func New(cfg Configurator, logger *log.Logger) *Application {
	return &Application{
		Config: cfg,
		Logs:   logger,
	}
}
