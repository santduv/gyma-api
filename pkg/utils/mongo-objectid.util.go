package utils

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StringToObjectID(id string) (primitive.ObjectID, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return primitive.NilObjectID, errors.New("invalid ObjectID")
	}

	return objectId, nil
}
