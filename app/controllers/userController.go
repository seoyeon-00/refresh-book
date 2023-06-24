package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// UserList godoc
// @Summary User list
// @Description Retrieve a list of users
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Hello user!"
// @Router /user [go get]
func UserList(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello user!")
}
