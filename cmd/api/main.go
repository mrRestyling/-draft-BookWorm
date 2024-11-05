package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"worm/internal/config"
	"worm/internal/handlers"
	"worm/internal/myApp"
	"worm/internal/server"
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
	// указанный в структуре конфигурации

	srv := server.New(fmt.Sprintf(":%d", cfg.Port), hand.Routes(), time.Minute, 10*time.Second, 30*time.Second)

	logger.Printf("start %s server on %s", cfg.Env, srv.Addr)

	err := srv.ListenAndServe()

	logger.Fatal(err)

}
