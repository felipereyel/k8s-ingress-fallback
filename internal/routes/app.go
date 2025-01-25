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
	app.Get("/_deployments/:namespace/:deployment", bind(svcs, details))
	app.Post("/_deployments/:namespace/:deployment", bind(svcs, toggle))

	app.Use("/_statics", staticsHandler)
	app.Use("/_healthz", healthzHandler)
	app.Use(notFoundHandler)

	return app
}
