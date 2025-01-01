package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santduv/gyma-api/internal/modules/auth/app/dto"
	app "github.com/santduv/gyma-api/internal/modules/auth/app/use-cases"
)

type AuthHandler struct {
	loginUseCase *app.LoginUseCase
}

func NewAuthHandler(loginUseCase *app.LoginUseCase) *AuthHandler {
	return &AuthHandler{
		loginUseCase: loginUseCase,
	}
}

func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	var loginDto dto.LoginDto

	if err := ctx.BodyParser(&loginDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result, errUseCase := h.loginUseCase.Execute(ctx.Context(), loginDto)

	if errUseCase != nil {
		return ctx.Status(errUseCase.StatusCode).JSON(errUseCase)
	}

	return ctx.Status(result.Status).JSON(result)
}
