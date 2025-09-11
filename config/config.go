package config

import (
	"log/slog"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Environment string

const (
	Dev  Environment = "dev"
	Prod Environment = "prod"
)

type Config struct {
	Environment   Environment
	Host          string
	Port          string
	LogLevel      slog.Level
	SessionSecret string
}

var (
	Global *Config
	once   sync.Once
)

func init() {
	once.Do(func() {
		Global = Load()
	})
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func loadBase() *Config {
	godotenv.Load()

	return &Config{
		Host: getEnv("HOST", "0.0.0.0"),
		Port: getEnv("PORT", "8080"),
		LogLevel: func() slog.Level {
			switch os.Getenv("LOG_LEVEL") {
			case "DEBUG":
				return slog.LevelDebug
			case "INFO":
				return slog.LevelInfo
			case "WARN":
				return slog.LevelWarn
			case "ERROR":
				return slog.LevelError
			default:
				return slog.LevelInfo
			}
		}(),
		SessionSecret: getEnv("SESSION_SECRET", "session-secret"),
	}
}
