package routes

import (
	"scaler/internal/services"

	"github.com/gofiber/fiber/v2"
)

func GetApp(svcs *services.Services) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	app.Post("/toggle/:namespace/:deployment", bind(svcs, deploymentToggle))
	app.Get("/", bind(svcs, deploymentList))

	app.Use("/healthz", healthzHandler)
	app.Use(notFoundHandler)

	return app
}
