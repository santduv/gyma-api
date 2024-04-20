package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	SuperRoutes "github.com/santduv/gyma-api/internal/super/interfaces/handler/routes"
)

func CreateApp() *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS, PATCH, HEAD",
	}))

	app.Use(requestid.New())

	router := app.Group("/api")

	// Register routes
	SuperRoutes.SetupRoutes(router.Group("/v1/supers"))

	return app
}
