package ports

import (
	"context"

	"github.com/santduv/gyma-api/internal/modules/gyms/domain/entities"
)

type FindGym struct {
	ID       *string
	Nickname *string
}

type GymRepository interface {
	Create(ctx context.Context, gym *entities.Gym) error
	FindOne(ctx context.Context, findGym *FindGym) (*entities.Gym, error)
}
