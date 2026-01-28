package config

import "os"

type Config struct {
	Env  string
	Port string
}

func Load() Config {
	return Config{
		Env:  getEnv("APP_ENV", "dev"),
		Port: getEnv("APP_PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
