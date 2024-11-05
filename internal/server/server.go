package server

import (
	"net/http"
	"time"
)

func New(addr string, handler http.Handler, IdleTimeout, ReadTimeout, WriteTimeout time.Duration) *http.Server {
	return &http.Server{
		Addr:         addr,
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}
