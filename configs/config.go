package configs

import (
	"os"
	"strconv"
	"time"
	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {

	// 서버 설정 정의
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}
}