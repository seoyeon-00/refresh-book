package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func FiberMiddleware(a *fiber.App) {
	a.Use(

		// 각 라우트에 CORS 추가하기
		cors.New(),
		
		// 로거 추가하기
		logger.New(),
	)
}