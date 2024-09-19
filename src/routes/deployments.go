package routes

import (
	"fmt"
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
	namespace := c.Params("namespace")
	deployment := c.Params("deployment")

	fmt.Println("Toggling deployment", namespace, deployment)

	return sendPage(c, components.NotFoundPage())
}
