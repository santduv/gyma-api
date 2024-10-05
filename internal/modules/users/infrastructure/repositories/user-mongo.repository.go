package repositories

import (
	"context"
	"time"

	"github.com/santduv/gyma-api/internal/modules/users/domain/entities"
	"github.com/santduv/gyma-api/internal/modules/users/domain/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoRepository struct {
	collection *mongo.Collection
}

func NewUserMongoRepository(collection *mongo.Collection) ports.UserRepository {
	return &UserMongoRepository{
		collection: collection,
	}
}

func (r *UserMongoRepository) Create(ctx context.Context, user *entities.User) error {
	_, err := r.collection.InsertOne(ctx, user)

	return err
}

func (r *UserMongoRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*entities.User, error) {
	var user entities.User

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserMongoRepository) Update(ctx context.Context, user *entities.User) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})

	return err
}

func (r *UserMongoRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"deletedAt": time.Now()}})

	return err
}
