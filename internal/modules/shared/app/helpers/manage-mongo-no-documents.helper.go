package helpers

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

func ManageMongoNoDocumentsError(err error) error {
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil
	}

	return err
}
