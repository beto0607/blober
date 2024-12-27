package core

import (
	"context"
	"log"

	"beto0607.com/blober/src/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DB_NAME = "blober"

var DBClient *mongo.Client
var DBHandler *mongo.Database

func ConnectToDB() {
	uri, err := config.GetEnvVar("MONGODB_URI")

	if err != nil {
		log.Panic("Couldn't get DB connection string")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))

	if err != nil {
		log.Panic("Couldn't create DB client")
	}

	DBClient = client

	dbHandler := client.Database(DB_NAME)
	DBHandler = dbHandler
	log.Println("DB connected :D")
}

func DisconnectDB() {
	if err := DBClient.Disconnect(context.TODO()); err != nil {
		log.Panic("Couldn't disconnectDB client")
	}
	log.Println("DB disconnected D:")
}
