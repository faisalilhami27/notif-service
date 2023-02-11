package usecases

import (
	"notification-service/broker/producer"
	"notification-service/src/repositories"
	"notification-service/src/usecases/v1/messageLog"
	"notification-service/src/usecases/v1/sms"
	"notification-service/src/usecases/v1/whatsapp"
)

type UseCase struct {
	MessageLogUseCase messageLog.MessageLogUseCase
	WhatsappUseCase   whatsapp.WhatsappUseCase
	SMSUseCase        sms.SMSUseCase
}

func NewUseCase(repository *repositories.Repository, producer *producer.Producer) *UseCase {
	messageLogUseCase := messageLog.NewMessageLogUseCase(repository.MessageLogRepository)
	whatsappUseCase := whatsapp.NewWhatsappUseCase(producer.WhatsappProducer)
	smsUseCase := sms.NewSMSUseCase(producer.SMSProducer)
	return &UseCase{
		MessageLogUseCase: messageLogUseCase,
		WhatsappUseCase:   whatsappUseCase,
		SMSUseCase:        smsUseCase,
	}
}
