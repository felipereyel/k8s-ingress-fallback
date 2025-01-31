package routes

import (
	"scaler/internal/config"
	"scaler/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func GetApp(svcs *services.Services, cfg config.ServerConfigs) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	if cfg.BasicPassword != "" {
		app.Use(basicauth.New(basicauth.Config{
			Users: map[string]string{
				"admin": cfg.BasicPassword,
			},
		}))
	}

	app.Get("/", bind(svcs, home))
	app.Get("/_deployments/:namespace/:deployment", bind(svcs, details))
	app.Post("/_deployments/:namespace/:deployment", bind(svcs, toggle))

	app.Use("/_statics", staticsHandler)
	app.Use("/_healthz", healthzHandler)
	app.Use(notFoundHandler)

	return app
}
