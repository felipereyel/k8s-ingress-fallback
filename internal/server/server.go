package server

import (
	"fallback/internal/config"
	"fallback/internal/routes"
	"fallback/internal/services"
)

func SetupAndListen() error {
	cfg := config.GetServerConfigs()

	svcs, err := services.Factory(cfg)
	if err != nil {
		return err
	}

	return routes.GetApp(svcs, cfg).Listen(cfg.ServerAddress)
}
