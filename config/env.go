package config

import (
	"os"

	"github.com/lpernett/godotenv"
)

type Config struct {
	Port       string
	PublicHost string
	DBUser     string
	DBPasswd   string
	DBHost     string
	DBPort     string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		Port:       getEnv("PORT", ":8080"),
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPasswd:   getEnv("DB_PASSWORD", "root"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBName:     getEnv("DB_NAME", "librarydb"),
		DBPort:     getEnv("DB_PORT", "3306"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
