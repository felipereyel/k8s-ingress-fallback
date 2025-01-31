package routes

import (
	"fallback/internal/config"
	"fallback/internal/services"

	"github.com/gofiber/fiber/v2"
)

func GetApp(svcs *services.Services, cfg config.ServerConfigs) *fiber.App {
	app := fiber.New(fiber.Config{ErrorHandler: errorHandler})

	app.Post("/_toggle", bind(svcs, toggle))
	app.Use("/_statics", staticsHandler)
	app.Use("/_healthz", healthzHandler)
	app.Use(bind(svcs, details))

	return app
}
