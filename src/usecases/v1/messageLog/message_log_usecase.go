package messageLog

import (
	"go.mongodb.org/mongo-driver/mongo"
	messageLogDT "notification-service/interfaces/dataTransfer/v1/messageLog"
	messageLogModel "notification-service/src/models/mongo/v1/messageLog"
	"notification-service/src/repositories/v1/messageLog"
)

type repository struct {
	repository messageLog.MessageLogRepository
}

type MessageLogUseCase interface {
	GetAllCategory() ([]messageLogModel.MessageLog, error)
	Create(input messageLogDT.MessageLogDT) (*mongo.InsertOneResult, error)
}

func NewMessageLogUseCase(categoryRepository messageLog.MessageLogRepository) *repository {
	return &repository{categoryRepository}
}

func (r *repository) GetAllCategory() ([]messageLogModel.MessageLog, error) {
	categories, err := r.repository.GetAllLog()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *repository) Create(input messageLogDT.MessageLogDT) (*mongo.InsertOneResult, error) {
	data := messageLogModel.MessageLog{
		MessageID:      input.MessageID,
		ReceiverNumber: input.ReceiverNumber,
		ErrorMessage:   input.ErrorMessage,
		MessageRaw:     input.MessageRaw,
		Provider:       input.Provider,
	}

	result, err := r.repository.CreateLog(data)
	if err != nil {
		return nil, err
	}

	return result, nil
}
