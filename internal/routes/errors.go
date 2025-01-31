package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func errorHandler(c *fiber.Ctx, err error) error {
	fmt.Printf("Route Error [%s]: %v\n", c.Path(), err)
	return c.SendStatus(fiber.StatusNotFound)
}
