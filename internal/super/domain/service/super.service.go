package service

import "github.com/gofiber/fiber/v2"

func WhoAmI() fiber.Map {
	return fiber.Map{
		"message": "You are super user (from service)",
	}
}
