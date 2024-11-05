package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"worm/internal/config"
	"worm/internal/handlers"
	"worm/internal/myApp"
)

func main() {

	// экземпляр структуры Config
	cfg := config.New()

	// Инициализируем логгер (обычный вывод; текущая дата и время)
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Объявляем экземпляр структуры приложения, которая содержит структуру конфигурации и логгер
	app := myApp.New(cfg, logger)

	hand := handlers.New(app)

	// Объявляем HTTP-сервер с настройками тайм-аута, который прослушивает порт,
	// указанный в структуре конфигурации, и использует созданный выше мультиплексор.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      hand.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("start %s server on %s", cfg.Env, srv.Addr)

	err := srv.ListenAndServe()
	logger.Fatal(err)

}
