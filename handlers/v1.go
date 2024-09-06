package handlers

import "github.com/gofiber/fiber/v2"

// V1Handler mengatur rute untuk versi 1
func V1Handler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"version": "v1",
	})
}
