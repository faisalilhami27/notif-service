package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"notification-service/config"
	"time"
)

type mongoClient struct {
	c *mongo.Client
}

var mongoInstance *mongoClient

func GetMongoClient() *mongo.Client {
	connection := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin",
		config.AppConfig.DBUsername,
		config.AppConfig.DBPassword,
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBDatabase,
	)

	client, err := mongo.NewClient(options.Client().ApplyURI(connection))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	mongoInstance = &mongoClient{
		c: client,
	}

	return mongoInstance.c
}
