package consumer

import (
	"notification-service/broker/consumer/sms"
	"notification-service/broker/consumer/whatsapp"
	"notification-service/src/usecases"
)

type Consumer struct {
	WhatsappConsumer whatsapp.WhatsappConsumer
	SMSConsumer      sms.SMSConsumer
}

func NewConsumer(kafkaBrokers []string, useCase usecases.UseCase) *Consumer {
	whatsappConsumer := whatsapp.NewWhatsappConsumer(kafkaBrokers, useCase.MessageLogUseCase)
	smsConsumer := sms.NewSMSConsumer(kafkaBrokers, useCase.MessageLogUseCase)
	return &Consumer{
		WhatsappConsumer: whatsappConsumer,
		SMSConsumer:      smsConsumer,
	}
}
