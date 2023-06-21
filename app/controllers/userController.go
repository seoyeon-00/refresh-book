package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func UserList(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello user!")
}