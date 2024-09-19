package routes

import (
	"scaler/src/components"
	"scaler/src/repositories/kluster"

	"github.com/gofiber/fiber/v2"
	v1 "k8s.io/api/apps/v1"
)

func deploymentList(c *fiber.Ctx) error {
	deployments, err := kluster.ListDeployments()
	if err != nil {
		deployments = []v1.Deployment{}
	}

	return sendPage(c, components.DeploymentListPage(deployments))
}

func deploymentToggle(c *fiber.Ctx) error {
	c.Set("HX-Redirect", "/")
	namespace := c.Params("namespace")
	deployment := c.Params("deployment")

	deployments, err := kluster.ListDeployments()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	for _, d := range deployments {
		if d.Namespace == namespace && d.Name == deployment {
			replicas := 0
			if d.Spec.Replicas != nil && *d.Spec.Replicas == 0 {
				replicas = 1
			}

			if err = kluster.ScaleDeployment(d, int32(replicas)); err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			}
		}
	}

	return c.SendStatus(fiber.StatusOK)
}
