package controllers

import (
	"notification-service/src/controllers/v1/messageLog"
	"notification-service/src/controllers/v1/sms"
	"notification-service/src/controllers/v1/whatsapp"
	"notification-service/src/usecases"
)

type Controller struct {
	MessageLogController messageLog.MessageLogController
	WhatsappController   whatsapp.WhatsappController
	SMSController        sms.SMSController
}

func NewController(useCase *usecases.UseCase) *Controller {
	messageLogController := messageLog.NewMessageLogController(useCase.MessageLogUseCase)
	whatsappController := whatsapp.NewWhatsappController(useCase.WhatsappUseCase)
	smsController := sms.NewSMSController(useCase.SMSUseCase)
	return &Controller{
		MessageLogController: messageLogController,
		WhatsappController:   whatsappController,
		SMSController:        smsController,
	}
}
