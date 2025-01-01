package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/santduv/gyma-api/internal/database"
	authRoutes "github.com/santduv/gyma-api/internal/modules/auth/infrastructure/routes"
	"github.com/santduv/gyma-api/internal/modules/shared/app/constants"
	httpErrors "github.com/santduv/gyma-api/internal/modules/shared/app/http-errors"
	"github.com/santduv/gyma-api/internal/modules/shared/app/types"
	userRoutes "github.com/santduv/gyma-api/internal/modules/users/infrastructure/routes"
)

func errorHandler(c *fiber.Ctx) error {
	err := c.Next()

	if err == nil {
		return nil
	}

	if httpErr, ok := err.(*httpErrors.HttpError); ok {
		return c.Status(httpErr.StatusCode).JSON(httpErr)
	}

	return c.Status(constants.HTTP_STATUS_INTERNAL_SERVER_ERROR).JSON(&httpErrors.HttpError{
		StatusCode: constants.HTTP_STATUS_INTERNAL_SERVER_ERROR,
		Message:    "Internal Server Error",
		Details: &types.JsonMap{
			"error": err.Error(),
		},
	})
}

func CreateApp() *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS, PATCH, HEAD",
	}))

	app.Use(errorHandler)
	app.Use(limiter.New())
	app.Use(requestid.New())

	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	router := app.Group("/api")

	// Connect to MongoDB
	database.ConnectToMongo()

	// Register routes
	userRoutes.SetupRoutes(router.Group("/v1/users"))
	authRoutes.SetupRoutes(router.Group("/v1/auth"))

	return app
}
