package serve

import (
	"scaler/src/config"
	"scaler/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Serve() error {
	app := fiber.New(fiber.Config{
		ErrorHandler: routes.ErrorHandler,
	})

	app.Use(cors.New())
	routes.Init(app)

	cfg := config.GetServerConfigs()
	return app.Listen(cfg.ServerAddress)
}
