package app

import (
	"context"
	"errors"

	"github.com/santduv/gyma-api/internal/modules/shared/app/helpers"
	httpErrors "github.com/santduv/gyma-api/internal/modules/shared/app/http-errors"
	"github.com/santduv/gyma-api/internal/modules/shared/app/types"
	"github.com/santduv/gyma-api/internal/modules/users/app/dto"
	"github.com/santduv/gyma-api/internal/modules/users/app/factories"
	"github.com/santduv/gyma-api/internal/modules/users/domain/ports"
	"github.com/santduv/gyma-api/pkg/utils"
)

type CreateUserUseCase struct {
	userRepository ports.UserRepository
	userFactory    *factories.UserFactory
}

func NewCreateUserUseCase(userRepository ports.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepository,
		userFactory:    &factories.UserFactory{},
	}
}

func (u *CreateUserUseCase) Execute(ctx context.Context, dto *dto.CreateUserDto) (*types.ApiResponse, *httpErrors.HttpError) {
	hashedPassword, err := utils.HashPassword(dto.Password)

	if err != nil {
		return nil, httpErrors.NewInternalServerError("failed to hash password", &types.JsonMap{
			"error": err.Error(),
		})
	}

	err = u.validateEmailExists(dto.Email)

	if err != nil {
		return nil, httpErrors.NewConflictError("email already exists")
	}

	dto.Password = hashedPassword
	user := u.userFactory.NewUserEntityFromDto(dto)

	err = u.userRepository.Create(ctx, user)

	if err != nil {
		return nil, httpErrors.NewInternalServerError("failed to create user", &types.JsonMap{
			"error": err.Error(),
		})
	}

	res := helpers.CreatedResponse("user created", u.userFactory.UserDtoFromEntity(user))

	return res, nil
}

func (u *CreateUserUseCase) validateEmailExists(email string) error {
	user, err := u.userRepository.FindOne(context.Background(), &ports.FindUser{
		Email: &email,
	})

	if err != nil {
		return helpers.ManageMongoNoDocumentsError(err)
	}

	if user != nil {
		return errors.New("email already exists")
	}

	return nil
}
