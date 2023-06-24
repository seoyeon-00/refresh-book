package routes

import (
	"refresh-bookstore/app/controllers"

	"github.com/gofiber/fiber/v2"
);

func Setup (app *fiber.App) {

	app.Get("/user", controllers.UserList)

}