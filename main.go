package main

import (
	"log"
	"refresh-bookstore/configs"
	"refresh-bookstore/database"
	_ "refresh-bookstore/docs" // docs 폴더를 import 합니다
	"refresh-bookstore/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger" // swag를 위한 fiber middleware
	"github.com/joho/godotenv"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for using swagger with fiber.
// @host localhost:3000
// @BasePath /
func main() {
	// .env 파일 로드
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := configs.LoadConfig()

	// Database 연결
	dbClient, err := database.PGConnection(config)
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	defer dbClient.Close()

	// Fiber 앱 생성
	app := fiber.New(config.FiberConfig)

	// Swagger 설정
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "http://127.0.0.1:3000/swagger/doc.json",
		DeepLinking: false,
		DocExpansion: "None",
	}))


	// 라우트 설정
	routes.Setup(app)

	// 서버 시작
	log.Fatal(app.Listen(":80"))
}
