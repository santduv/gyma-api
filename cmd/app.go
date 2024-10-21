package cmd

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/santduv/gyma-api/internal/database"
	"github.com/santduv/gyma-api/internal/modules/shared/app/constants"
	httpErrors "github.com/santduv/gyma-api/internal/modules/shared/app/http-errors"
	"github.com/santduv/gyma-api/internal/modules/shared/app/types"
	UserRoutes "github.com/santduv/gyma-api/internal/modules/users/infrastructure/routes"
)

func errorHandler(c *fiber.Ctx) error {
	err := c.Next()

	if err == nil {
		return nil
	}

	fmt.Println(err)

	if errors.Is(err, &httpErrors.HttpError{}) {
		errProps := err.(*httpErrors.HttpError)

		return c.Status(errProps.StatusCode).JSON(errProps)
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

	app.Use(requestid.New())

	router := app.Group("/api")

	// Connect to MongoDB
	database.ConnectToMongo()

	// Register routes
	UserRoutes.SetupRoutes(router.Group("/v1/users"))

	return app
}
