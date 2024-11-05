package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Версия приложения
// Позже мы будем генерировать версию автоматически
const Version = "1.0.0"

type Handlers struct {
	APP AppInt
}

type AppInt interface {
}

func New(app AppInt) *Handlers {
	return &Handlers{
		APP: app,
	}
}

func (h *Handlers) Routes() *chi.Mux {

	// Инициализируем маршрутизатор chi
	route := chi.NewRouter()

	route.Get("/v1/healthcheck", h.HealhcheckHandler)
	route.Post("/v1/books", h.CreateHandler)
	route.Get("/v1/books/{id}", h.ShowBookHandler)

	return route

}

func (h *Handlers) HealhcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Для проверки
	log.Println("запрос на HealhcheckHandler")
	fmt.Fprintln(w, "status: available")
	// fmt.Fprintf(w, "environment: %s\n", h.APP.Configurator.Env)
	fmt.Fprintf(w, "version: %s\n", Version)

}
