package database

import (
	"context"
	"log"

	"github.com/santduv/gyma-api/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func ConnectToMongo() {
	uri := config.Envs.MongoURI
	dbName := config.Envs.MongoDB

	context := context.Background()

	client, err := mongo.Connect(context, options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Database(dbName).Client().Ping(context, nil)

	if err != nil {
		log.Fatal(err)
	}

	db = client.Database(dbName)
	log.Println("Connected to MongoDB")
}

func GetCollection(collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}
