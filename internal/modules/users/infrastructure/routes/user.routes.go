package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santduv/gyma-api/internal/modules/users/infrastructure"
)

func SetupRoutes(router fiber.Router) {
	userCtx := infrastructure.NewUserContext()

	router.Post("/", userCtx.UserHandler.CreateUser)
}
