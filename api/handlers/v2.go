package handlers

import "github.com/gofiber/fiber/v2"

// V2Handler mengatur rute untuk versi 2
func V2Handler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"version": "v2",
	})
}
