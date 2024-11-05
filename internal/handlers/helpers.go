package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func ReadID(r *http.Request) (int64, error) {
	param := chi.URLParam(r, "id")

	log.Println("param: ", param)

	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil || id < 0 {
		return 0, errors.New("invalid parametr")
	}

	return id, nil
}
