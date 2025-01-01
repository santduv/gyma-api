package repositories

import (
	"context"

	"github.com/santduv/gyma-api/internal/modules/gyms/domain/entities"
	"github.com/santduv/gyma-api/internal/modules/gyms/domain/ports"
	"github.com/santduv/gyma-api/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type GymMongoRepository struct {
	collection *mongo.Collection
}

func NewGymMongoRepository(collection *mongo.Collection) ports.GymRepository {
	return &GymMongoRepository{
		collection: collection,
	}
}

func (r *GymMongoRepository) Create(ctx context.Context, gym *entities.Gym) error {
	_, err := r.collection.InsertOne(ctx, gym)

	return err
}

func (r *GymMongoRepository) FindOne(ctx context.Context, findGym *ports.FindGym) (*entities.Gym, error) {
	var gym entities.Gym

	filter := map[string]interface{}{}

	if findGym.ID != nil {
		gymObjectId, err := utils.StringToObjectID(*findGym.ID)

		if err != nil {
			return nil, err
		}

		filter["_id"] = gymObjectId
	}

	if findGym.Nickname != nil {
		filter["nickname"] = *findGym.Nickname
	}

	err := r.collection.FindOne(ctx, filter).Decode(&gym)

	if err != nil {
		return nil, err
	}

	return &gym, nil
}
