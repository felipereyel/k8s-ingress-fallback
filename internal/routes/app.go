package routes

import (
	"scaler/internal/services"

	"github.com/gofiber/fiber/v2"
)

func GetApp(svcs *services.Services) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	app.Get("/", bind(svcs, home))
	app.Get("/:namespace/:deployment", bind(svcs, details))
	app.Post("/_toggle/:namespace/:deployment", bind(svcs, toggle))

	app.Use("/_healthz", healthzHandler)
	app.Use(notFoundHandler)

	return app
}
