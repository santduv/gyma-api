package factories

import (
	"time"

	"github.com/santduv/gyma-api/internal/modules/users/app/dto"
	"github.com/santduv/gyma-api/internal/modules/users/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserFactory struct{}

func (f *UserFactory) NewUserEntityFromDto(dto *dto.CreateUserDto) *entities.User {
	now := time.Now()

	return &entities.User{
		ID:        primitive.NewObjectID(),
		Email:     dto.Email,
		Nickname:  dto.Nickname,
		Password:  dto.Password,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: nil,
	}
}

func (f *UserFactory) UserDtoFromEntity(user *entities.User) *dto.UserDto {
	return &dto.UserDto{
		ID:        user.ID.Hex(),
		Email:     user.Email,
		Nickname:  user.Nickname,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
}
