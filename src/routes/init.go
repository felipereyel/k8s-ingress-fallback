package routes

import (
	"github.com/gofiber/fiber/v2"
)

func healthzHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func Init(app *fiber.App) error {

	app.Post("/toggle/:id", deploymentToggle)
	app.Get("/", deploymentList)

	app.Use("/healthz", healthzHandler)
	app.Use(notFoundHandler)

	return nil
}
