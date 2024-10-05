package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santduv/gyma-api/internal/modules/users/app/dto"
	app "github.com/santduv/gyma-api/internal/modules/users/app/use-cases"
)

type UserHandler struct {
	createUserUseCase *app.CreateUserUseCase
}

func NewUserHandler(createUserUseCase *app.CreateUserUseCase) *UserHandler {
	return &UserHandler{
		createUserUseCase: createUserUseCase,
	}
}

func (h *UserHandler) CreateUser(ctx *fiber.Ctx) error {
	var createUserDto dto.CreateUserDto

	if err := ctx.BodyParser(&createUserDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user, err := h.createUserUseCase.Execute(ctx.Context(), &createUserDto)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(user)
}
