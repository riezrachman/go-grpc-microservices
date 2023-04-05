package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ListenAddress      string   `config:"LISTEN_ADDRESS"`
	CorsAllowedHeaders []string `config:"CORS_ALLOWED_HEADERS"`
	CorsAllowedMethods []string `config:"CORS_ALLOWED_METHODS"`
	CorsAllowedOrigins []string `config:"CORS_ALLOWED_ORIGINS"`
	JWTSecret          string   `config:"JWT_SECRET"`
	JWTDuration        string   `config:"JWT_DURATION"`
	DSN                string   `config:"DB_DSN"`
}

var config *Config

func initConfig() {

	godotenv.Load(".env")

	config = &Config{
		ListenAddress: fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")),
		CorsAllowedHeaders: []string{
			"Connection", "User-Agent", "Referer",
			"Accept", "Accept-Language", "Content-Type",
			"Content-Language", "Content-Disposition", "Origin",
			"Content-Length", "Authorization", "ResponseType",
			"X-Requested-With", "X-Forwarded-For", "X-Auth-Role",
		},
		CorsAllowedMethods: []string{"GET", "POST", "DELETE"},
		CorsAllowedOrigins: []string{"*"},
		JWTSecret:          getEnv("JWT_SECRET", "secret"),
		JWTDuration:        getEnv("JWT_DURATION", "30m"),
		DSN:                getEnv("DB_DSN", ""),
	}

}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
