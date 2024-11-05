package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) Routes() *chi.Mux {

	// Инициализируем маршрутизатор chi
	router := chi.NewRouter()

	router.Get("/v1/healthcheck", a.HealhcheckHandler)
	router.Post("/v1/books", a.CreateHandler)
	router.Get("/v1/books/{id}", a.ShowBookHandler)

	return router

}

func (a *application) HealhcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Для проверки
	log.Println("запрос на HealhcheckHandler")
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", a.config.env)
	fmt.Fprintf(w, "version: %s\n", version)

}
