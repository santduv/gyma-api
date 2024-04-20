package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santduv/gyma-api/internal/super/interfaces/handler/controller"
)

func SetupRoutes(router fiber.Router) {
	router.Get("/whoami", controller.WhoAmI)
}
