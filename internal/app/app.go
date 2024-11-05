package app

import (
	"fmt"
	"log"
	"net/http"
	"worm/internal/config"

	"github.com/go-chi/chi"
)

// Версия приложения
// Позже мы будем генерировать версию автоматически
const Version = "1.0.0"

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

func (a *Application) Routes() *chi.Mux {

	// Инициализируем маршрутизатор chi
	router := chi.NewRouter()

	router.Get("/v1/healthcheck", a.HealhcheckHandler)
	router.Post("/v1/books", a.CreateHandler)
	router.Get("/v1/books/{id}", a.ShowBookHandler)

	return router

}

func (a *Application) HealhcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Для проверки
	log.Println("запрос на HealhcheckHandler")
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", a.Config.Env)
	fmt.Fprintf(w, "version: %s\n", Version)

}
