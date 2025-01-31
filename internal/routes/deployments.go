package routes

import (
	"fallback/internal/components"
	"fallback/internal/services"

	"github.com/gofiber/fiber/v2"
)

func details(svcs *services.Services, c *fiber.Ctx) error {
	hostname := c.Hostname()

	d, err := svcs.KubeClient.GetDeploymentForHostname(hostname)
	if err != nil {
		return err
	}

	return sendPage(c, components.DetailsPage(d))
}

func toggle(svcs *services.Services, c *fiber.Ctx) error {
	c.Set("HX-Refresh", "true")
	hostname := c.Hostname()

	d, err := svcs.KubeClient.GetDeploymentForHostname(hostname)
	if err != nil {
		return err
	}

	replicas := 0
	if d.Spec.Replicas != nil && *d.Spec.Replicas == 0 {
		replicas = 1
	}

	if err = svcs.KubeClient.ScaleDeployment(d, int32(replicas)); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
