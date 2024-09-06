package handlers

import "github.com/gofiber/fiber/v2"

// RootHandler mengatur rute root
func RootHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"uri":  ctx.Request().URI().String(),
		"path": ctx.Path(),
	})
}
