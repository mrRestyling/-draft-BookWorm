package config

import (
	"flag"
)

// все параметры конфигурации для нашего приложения
type Config struct {
	Port int
	Env  string
}

func (c *Config) ParseFlags() {
	flag.IntVar(&c.Port, "port", 8080, "API server port")
	flag.StringVar(&c.Env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()
}

func New() *Config {
	c := &Config{}
	c.ParseFlags()
	return c
}
