package app

import (
	"context"
	"errors"

	"github.com/santduv/gyma-api/internal/modules/gyms/app/dto"
	"github.com/santduv/gyma-api/internal/modules/gyms/app/factories"
	"github.com/santduv/gyma-api/internal/modules/gyms/domain/ports"
	"github.com/santduv/gyma-api/internal/modules/shared/app/helpers"
	httpErrors "github.com/santduv/gyma-api/internal/modules/shared/app/http-errors"
	"github.com/santduv/gyma-api/internal/modules/shared/app/types"
)

type CreateGymUseCase struct {
	gymRepository ports.GymRepository
	gymFactory    *factories.GymFactory
}

func NewCreateGymUseCase(gymRepository ports.GymRepository) *CreateGymUseCase {
	return &CreateGymUseCase{
		gymRepository: gymRepository,
		gymFactory:    &factories.GymFactory{},
	}
}

func (u *CreateGymUseCase) Execute(ctx context.Context, createGymDto *dto.CreateGymDto) (*types.ApiResponse, *httpErrors.HttpError) {
	err := u.validateDto(createGymDto)

	if err != nil {
		return nil, httpErrors.NewBadRequestError("invalid request body", &types.JsonMap{
			"error": err.Error(),
		})
	}

	err = u.validateNicknameExists(ctx, createGymDto.Nickname)

	if err != nil {
		return nil, httpErrors.NewConflictError("nickname already exists", nil)
	}

	gymEntity, err := u.gymFactory.NewGymEntityFromDto(createGymDto)

	if err != nil {
		return nil, httpErrors.NewInternalServerError("failed to create gym", &types.JsonMap{
			"error": err.Error(),
		})
	}

	err = u.gymRepository.Create(ctx, gymEntity)

	if err != nil {
		return nil, httpErrors.NewInternalServerError("failed to create gym", &types.JsonMap{
			"error": err.Error(),
		})
	}

	gymDto := u.gymFactory.NewGymDtoFromEntity(gymEntity)

	res := helpers.CreatedResponse("gym created", gymDto)

	return res, nil
}

func (u *CreateGymUseCase) validateNicknameExists(ctx context.Context, nickname string) error {
	gym, err := u.gymRepository.FindOne(ctx, &ports.FindGym{
		Nickname: &nickname,
	})

	if err != nil {
		return helpers.ManageMongoNoDocumentsError(err)
	}

	if gym != nil {
		return errors.New("nickname already exists")
	}

	return nil
}

func (u *CreateGymUseCase) validateDto(createGymDto *dto.CreateGymDto) error {
	if createGymDto.Name == "" {
		return errors.New("name is required")
	}

	if createGymDto.Nickname == "" {
		return errors.New("nickname is required")
	}

	return nil
}
