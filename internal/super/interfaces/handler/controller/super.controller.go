package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santduv/gyma-api/internal/super/domain/service"
)

// WhoAmI Provides information about the current user
func WhoAmI(c *fiber.Ctx) error {
	result := service.WhoAmI()

	return c.JSON(result)
}
