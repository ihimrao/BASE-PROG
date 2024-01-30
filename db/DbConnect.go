package database

import (
	"context"
	middlewares "go-base-fs/handlers"
	"log"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func DatabaseConnection() *mongo.Client {
	options := options.Client().ApplyURI(middlewares.GetEnvVar("MONGO_URI"))
	client, err := mongo.Connect(context.TODO(), options)

	if err != nil {
		log.Fatal("Error connecting to MongoDB", middlewares.GetEnvVar("MONGO_URI"))
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Error connecting to MongoDB")
	}
	color.Cyan("MongoDB Connected Successfully")
	return client
}
