package config

import (
	"flag"
)

// все параметры конфигурации для нашего приложения
type Config struct {
	Port int
	Env  string
}

func New() *Config {

	var port int
	var env string
	// Записываем значения флагов командной строки port и env в структуру конфигурации
	flag.IntVar(&port, "port", 8000, "API server port")
	flag.StringVar(&env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	return &Config{
		Port: port,
		Env:  env,
	}
}
