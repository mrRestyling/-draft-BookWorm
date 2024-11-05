package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// Обработчик для эндпоинта POST
func (h *Handlers) CreateHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("запрос на CreateHandler")

	log.Println("ДОСТУП ЕСТЬ")

	fmt.Fprint(w, "Create new book")
}

// Обработчик для эндпоинта GET
func (h *Handlers) ShowBookHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("запрос на ShowBookHandler")

	// Получаем id из URL вспомогательным методом
	id, err := ReadID(r)
	if err != nil {
		log.Printf("ошибка при получении id: %v", err)
		http.NotFound(w, r)
	}

	fmt.Fprintf(w, "details of book %d\n", id)

}
