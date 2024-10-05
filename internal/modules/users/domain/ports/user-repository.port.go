package ports

import (
	"context"

	"github.com/santduv/gyma-api/internal/modules/users/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	FindByID(ctx context.Context, id primitive.ObjectID) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}
