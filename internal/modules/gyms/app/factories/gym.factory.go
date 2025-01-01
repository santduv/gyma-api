package factories

import (
	"time"

	"github.com/santduv/gyma-api/internal/modules/gyms/app/dto"
	"github.com/santduv/gyma-api/internal/modules/gyms/domain/entities"
	"github.com/santduv/gyma-api/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GymFactory struct{}

func (f *GymFactory) NewGymEntityFromDto(dto *dto.CreateGymDto) (*entities.Gym, error) {
	now := time.Now()

	userObjectId, err := utils.StringToObjectID(dto.CreatedBy)

	if err != nil {
		return nil, err
	}

	entity := &entities.Gym{
		ID:        primitive.NewObjectID(),
		Name:      dto.Name,
		Nickname:  dto.Nickname,
		Active:    true,
		CreatedBy: userObjectId,
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: nil,
	}

	return entity, nil
}

func (f *GymFactory) NewGymDtoFromEntity(gym *entities.Gym) *dto.GymDto {
	var deletedAt *string

	if gym.DeletedAt != nil {
		formattedDate := gym.DeletedAt.Format(time.RFC3339)
		deletedAt = &formattedDate
	}

	return &dto.GymDto{
		ID:        gym.ID.Hex(),
		Name:      gym.Name,
		Nickname:  gym.Nickname,
		Active:    gym.Active,
		CreatedBy: gym.CreatedBy.Hex(),
		CreatedAt: gym.CreatedAt.Format(time.RFC3339),
		UpdatedAt: gym.UpdatedAt.Format(time.RFC3339),
		DeletedAt: deletedAt,
	}
}
