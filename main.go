package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	producer2 "notification-service/broker/producer"
	"notification-service/config"
	"notification-service/database"
	"notification-service/routes"
	"notification-service/src/controllers"
	"notification-service/src/repositories"
	"notification-service/src/usecases"
	"notification-service/utils"
	"time"
)

func init() {
	err := utils.LoadEnv()
	if err != nil {
		return
	}
	config.InitApp()
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Now().In(loc)
}

func main() {
	db := database.GetMongoClient()
	defer func(db *mongo.Client, ctx context.Context) {
		err := db.Disconnect(ctx)
		if err != nil {
			panic(err)
		}
	}(db, context.Background())

	// load kafka producer
	addresses := config.AppConfig.KafkaBrokerAddress
	producer := producer2.NewProducer(addresses)

	repository := repositories.NewRepository(db)
	useCase := usecases.NewUseCase(repository, producer)
	controller := controllers.NewController(useCase)
	router := routes.NewRoute(*controller)
	router.Run()
}
