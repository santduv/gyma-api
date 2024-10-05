package app

import (
	"context"

	"github.com/santduv/gyma-api/internal/modules/users/app/dto"
	"github.com/santduv/gyma-api/internal/modules/users/app/factories"
	"github.com/santduv/gyma-api/internal/modules/users/domain/entities"
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

func (u *CreateUserUseCase) Execute(ctx context.Context, dto *dto.CreateUserDto) (*entities.User, error) {
	hashedPassword, err := utils.HashPassword(dto.Password)

	if err != nil {
		panic(err)
	}

	dto.Password = hashedPassword
	user := u.userFactory.NewUserEntityFromDto(dto)

	err = u.userRepository.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
