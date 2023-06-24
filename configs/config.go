package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	DatabaseURL            string
	FiberConfig            fiber.Config
	DBMaxConnections       int
	DBMaxIdleConnections   int
	DBMaxLifetimeConnections int
}

func LoadConfig() Config {
	maxConn, _ := strconv.Atoi(getEnv("DB_MAX_CONNECTIONS", "10"))
	maxIdleConn, _ := strconv.Atoi(getEnv("DB_MAX_IDLE_CONNECTIONS", "2"))
	maxLifetimeConn, _ := strconv.Atoi(getEnv("DB_MAX_LIFETIME_CONNECTIONS", "60"))

	return Config{
		DatabaseURL:              getEnv("DATABASE_URL", ""),
		FiberConfig:              getFiberConfig(),
		DBMaxConnections:         maxConn,
		DBMaxIdleConnections:     maxIdleConn,
		DBMaxLifetimeConnections: maxLifetimeConn,
	}
}

func getFiberConfig() fiber.Config {
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
