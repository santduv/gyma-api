package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Gym struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Nickname  string             `json:"nickname" bson:"nickname"`
	Active    bool               `json:"active" bson:"active"`
	CreatedBy primitive.ObjectID `json:"createdBy" bson:"createdBy"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	DeletedAt *time.Time         `json:"deletedAt" bson:"deletedAt"`
}
