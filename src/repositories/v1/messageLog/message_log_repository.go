package messageLog

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	models "notification-service/src/models/mongo/v1/messageLog"
	"notification-service/utils"
	"time"
)

type MessageLogRepository interface {
	CreateLog(messageLog models.MessageLog) (*mongo.InsertOneResult, error)
	GetAllLog() ([]models.MessageLog, error)
}

type repository struct {
	db *mongo.Client
}

func NewMessageLogRepository(db *mongo.Client) MessageLogRepository {
	return &repository{db}
}

func (r *repository) CreateLog(messageLog models.MessageLog) (*mongo.InsertOneResult, error) {
	location, _ := time.LoadLocation("Asia/Jakarta")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := r.db.Database("notification_service").Collection("message_logs")
	messageLog.CreatedAt = time.Now().In(location)
	result, err := collection.InsertOne(ctx, messageLog)
	if !utils.MongoErrorDatabaseException(err) {
		return nil, err
	}
	return result, nil
}

func (r *repository) GetAllLog() ([]models.MessageLog, error) {
	var result []models.MessageLog
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := r.db.Database("notification_service").Collection("message_logs")
	cursor, err := collection.Find(ctx, bson.M{})

	if !utils.MongoErrorDatabaseException(err) {
		return nil, err
	}

	for cursor.Next(ctx) {
		var data models.MessageLog
		err := cursor.Decode(&data)
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}

	return result, nil
}
