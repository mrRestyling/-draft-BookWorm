package main

import (
	"fmt"
	"net/http"
)

// Обработчик для эндпоинта POST
func (a *application) CreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Create new book")
}

// Обработчик для эндпоинта GET
func (a *application) ShowBookHandler(w http.ResponseWriter, r *http.Request) {

	// Получаем id из URL вспомогательным методом
	id, err := a.ReadID(r)
	if err != nil {
		http.NotFound(w, r)
	}

	fmt.Fprintf(w, "detalis of book %d\n", id)

}
