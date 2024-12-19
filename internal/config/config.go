package config

import (
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	LogLevel string `envconfig:"LOG_LEVEL" default:"DEBUG"`
	Port     int    `envconfig:"PORT"`
}

func Get() *Config {
	once.Do(func() {
		conf = &Config{}
		godotenv.Load()
		envconfig.MustProcess("", conf)
	})

	return conf
}
