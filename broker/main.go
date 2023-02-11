package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"notification-service/broker/consumer"
	producer2 "notification-service/broker/producer"
	"notification-service/config"
	"notification-service/database"
	"notification-service/src/repositories"
	"notification-service/src/usecases"
	"notification-service/utils"
	"os"
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

	args := os.Args[1:]
	var consumerType string

	if len(args) > 0 {
		consumerType = args[0]
	}
	// load kafka consumer
	addresses := config.AppConfig.KafkaBrokerAddress
	repository := repositories.NewRepository(db)
	producer := producer2.NewProducer(addresses)
	useCase := usecases.NewUseCase(repository, producer)
	switch consumerType {
	case "whatsapp":
		whatsappConsumer := consumer.NewConsumer(addresses, *useCase)
		whatsappConsumer.WhatsappConsumer.ReadMessage(context.Background())
	case "sms":
		smsConsumer := consumer.NewConsumer(addresses, *useCase)
		smsConsumer.SMSConsumer.ReadMessage(context.Background())
	}
}
