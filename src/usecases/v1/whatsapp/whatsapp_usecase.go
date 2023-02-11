package whatsapp

import (
	"context"
	"github.com/google/uuid"
	"notification-service/broker/domain"
	whatsapp2 "notification-service/broker/producer/whatsapp"
	"notification-service/interfaces/dataTransfer/v1/whatsapp"
	"notification-service/utils"
)

type whatsappUseCase struct {
	producer whatsapp2.WhatsappProducer
}

func NewWhatsappUseCase(producer whatsapp2.WhatsappProducer) *whatsappUseCase {
	return &whatsappUseCase{
		producer: producer,
	}
}

type WhatsappUseCase interface {
	PublishMessage(ctx context.Context, data whatsapp.WhatsappBody) error
}

func (w *whatsappUseCase) PublishMessage(ctx context.Context, data whatsapp.WhatsappBody) error {
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
