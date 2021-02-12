package mongoDB

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/jethro91/fiber-mongo-redis/src/config"
)

var (
	Client *mongo.Client
)

func CreateMongoDBConnection() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.MONGO_URI))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	Client = client
}
