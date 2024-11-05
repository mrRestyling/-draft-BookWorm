package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"worm/internal/app"
	"worm/internal/config"
)

func main() {

	// экземпляр структуры Config
	cfg := config.New()

	// Инициализируем логгер (обычный вывод; текущая дата и время)
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Объявляем экземпляр структуры приложения, которая содержит структуру конфигурации и логгер

	app := app.New(cfg, logger)

	// Объявляем HTTP-сервер с настройками тайм-аута, который прослушивает порт,
	// указанный в структуре конфигурации, и использует созданный выше мультиплексор.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("start %s server on %s", cfg.Env, srv.Addr)

	err := srv.ListenAndServe()
	logger.Fatal(err)

}
