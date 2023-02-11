package producer

import (
	"notification-service/broker/producer/sms"
	"notification-service/broker/producer/whatsapp"
)

type Producer struct {
	WhatsappProducer whatsapp.WhatsappProducer
	SMSProducer      sms.SMSProducer
}

func NewProducer(kafkaBrokers []string) *Producer {
	whatsappProducer := whatsapp.NewWhatsappProducer(kafkaBrokers)
	smsProducer := sms.NewSMSProducer(kafkaBrokers)
	return &Producer{
		WhatsappProducer: whatsappProducer,
		SMSProducer:      smsProducer,
	}
}
