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
	LogLevel          string `envconfig:"LOG_LEVEL" default:"DEBUG"`
	Port              int    `envconfig:"PORT"`
	CORSAllowedOrigin string `envconfig:"CORS_ALLOWED_ORIGINS"`
	DBHost            string `envconfig:"DB_HOST"`
	DBPort            string `envconfig:"DB_PORT"`
	DBUser            string `envconfig:"DB_USER"`
	DBPassword        string `envconfig:"DB_PASSWORD"`
	DBName            string `envconfig:"DB_NAME"`
}

func Get() *Config {
	once.Do(func() {
		conf = &Config{}
		godotenv.Load()
		envconfig.MustProcess("", conf)
	})

	return conf
}
