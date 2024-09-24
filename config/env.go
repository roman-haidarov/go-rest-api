package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPasswd   string
	DBAddress  string
	DBName     string
	ApiKey     string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPasswd:   getEnv("DB_PASSWORD", "Password"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), "3306"),
		DBName:     getEnv("DB_NAME", "development"),
		ApiKey:     getEnv("API_KEY", "-"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
