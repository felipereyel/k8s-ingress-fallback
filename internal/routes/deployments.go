package routes

import (
	"scaler/internal/components"
	"scaler/internal/services"

	"github.com/gofiber/fiber/v2"
)

func deploymentList(svcs *services.Services, c *fiber.Ctx) error {
	deployments, err := svcs.KubeClient.ListDeployments()
	if err != nil {
		return err
	}

	return sendPage(c, components.DeploymentListPage(deployments))
}

func deploymentToggle(svcs *services.Services, c *fiber.Ctx) error {
	c.Set("HX-Redirect", "/")
	namespace := c.Params("namespace")
	deployment := c.Params("deployment")

	deployments, err := svcs.KubeClient.ListDeployments()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	for _, d := range deployments {
		if d.Namespace == namespace && d.Name == deployment {
			replicas := 0
			if d.Spec.Replicas != nil && *d.Spec.Replicas == 0 {
				replicas = 1
			}

			if err = svcs.KubeClient.ScaleDeployment(d, int32(replicas)); err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			}
		}
	}

	return c.SendStatus(fiber.StatusOK)
}
