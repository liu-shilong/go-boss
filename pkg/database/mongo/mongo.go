package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MongoClient *mongo.Client

func InitMongo() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	var err error
	conn, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	MongoClient = conn
}
