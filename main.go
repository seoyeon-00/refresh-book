package main

import (
	"github.com/gofiber/fiber/v2"
	"refresh-bookstore/database"
	"refresh-bookstore/routes"
)

func main() {
	app := fiber.New()

	database.MysqlConnection()
	routes.Setup(app)

  app.Listen(":3000")
}