package sms

import (
	"context"
	"github.com/google/uuid"
	"notification-service/broker/domain"
	"notification-service/broker/producer/sms"
	smsDT "notification-service/interfaces/dataTransfer/v1/sms"
	"notification-service/utils"
)

type smsUseCase struct {
	producer sms.SMSProducer
}

func NewSMSUseCase(producer sms.SMSProducer) *smsUseCase {
	return &smsUseCase{
		producer: producer,
	}
}

type SMSUseCase interface {
	PublishMessage(ctx context.Context, data smsDT.SMSBody) error
}

func (w *smsUseCase) PublishMessage(ctx context.Context, data smsDT.SMSBody) error {
	kafkaMessage := domain.KafkaMessage{
		ID:   uuid.NewString(),
		Data: data,
	}

	qmData, err := kafkaMessage.MarshalBinary()
	if err != nil {
		utils.GetErrorLog(err)
		return err
	}

	err = w.producer.WriteMessage(ctx, domain.QueueMessage{
		Key:   []byte("send_notification"),
		Value: qmData,
	})

	if err != nil {
		utils.GetErrorLog(err)
		return err
	}

	return nil
}
