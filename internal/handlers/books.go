package handlers

import (
	"fmt"
	"net/http"

	"github.com/ktr0731/evans/app"
)

type Handlers struct {
	APP *app.App
}

// Обработчик для эндпоинта POST
func (h *Handlers) CreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Create new book")
}

// Обработчик для эндпоинта GET
func (h *Handlers) ShowBookHandler(w http.ResponseWriter, r *http.Request) {

	// Получаем id из URL вспомогательным методом
	id, err := a.ReadID(r)
	if err != nil {
		http.NotFound(w, r)
	}

	fmt.Fprintf(w, "detalis of book %d\n", id)

}
