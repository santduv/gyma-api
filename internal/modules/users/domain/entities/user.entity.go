package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Email     string             `json:"email" bson:"email"`
	Nickname  string             `json:"nickname" bson:"nickname"`
	Password  string             `json:"password" bson:"password"`
	FirstName string             `json:"firstName" bson:"firstName"`
	LastName  string             `json:"lastName" bson:"lastName"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	DeletedAt *time.Time         `json:"deletedAt" bson:"deletedAt"`
}
