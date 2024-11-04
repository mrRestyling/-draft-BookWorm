package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Версия приложения
// Позже мы будем генерировать версию автоматически
const version = "1.0.0"

// все параметры конфигурации для нашего приложения
type Config struct {
	port int
	env  string
}

// Определим структуру приложения, которая будет содержать зависимости для
// обработчиков HTTP, вспомогательных функций и middleware
type application struct {
	config Config
	logs   *log.Logger
}

func main() {

	// экземпляр структуры Config
	var cfg Config

	// Записываем значения флагов командной строки port и env в структуру конфигурации
	flag.IntVar(&cfg.port, "port", 8000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Инициализируем логгер
	// 1. стандартный вывод
	// 2. текущая дата и время
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Объявляем экземпляр структуры приложения, которая содержит структуру конфигурации и логгер
	app := application{
		config: cfg,
		logs:   logger,
	}

	// Объявляем новый мультиплексор и добавляем маршрут `/v1/healthcheck`,
	// который будет перенаправлять запросы в метод `healhcheckHandler`

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.HealhcheckHandler)

	// http.HandleFunc("/v1/healthcheck", app.HealhcheckHandler)

	// Объявляем HTTP-сервер с настройками тайм-аута, который прослушивает порт,
	// указанный в структуре конфигурации, и использует созданный выше мультиплексор.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("start %s server on %s", cfg.env, srv.Addr)

	err := srv.ListenAndServe()
	logger.Fatal(err)

}

// аналог для терминала
// curl -i localhost:8000/v1/healthcheck

func (a *application) HealhcheckHandler(w http.ResponseWriter, r *http.Request) {
	// 2
	log.Println("запрос на HealhcheckHandler")
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", a.config.env)
	fmt.Fprintf(w, "version: %s\n", version)

}
