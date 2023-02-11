package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
	"notification-service/src/repositories/v1/messageLog"
)

type Repository struct {
	MessageLogRepository messageLog.MessageLogRepository
}

func NewRepository(db *mongo.Client) *Repository {
	categoryRepository := messageLog.NewMessageLogRepository(db)
	return &Repository{
		MessageLogRepository: categoryRepository,
	}
}
