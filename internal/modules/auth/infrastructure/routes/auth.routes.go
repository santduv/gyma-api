package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santduv/gyma-api/internal/modules/auth/infrastructure"
)

func SetupRoutes(router fiber.Router) {
	authCtx := infrastructure.NewAuthContext()

	router.Post("/login", authCtx.AuthHandler.Login)
}
