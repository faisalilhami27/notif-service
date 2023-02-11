package whatsapp

import (
	"context"
	"github.com/segmentio/kafka-go"
	"notification-service/broker/domain"
	"notification-service/config"
	"notification-service/utils"
)

type WhatsappProducerConfig struct {
	KafkaBroker []string
}

type WhatsappProducer interface {
	WriteMessage(ctx context.Context, data ...domain.QueueMessage) error
}

func NewWhatsappProducer(kafkaBroker []string) *WhatsappProducerConfig {
	return &WhatsappProducerConfig{
		KafkaBroker: kafkaBroker,
	}
}

func (p WhatsappProducerConfig) WriteMessage(ctx context.Context, data ...domain.QueueMessage) error {
	topic := config.AppConfig.KafkaWATopic
	kafkaWriter := utils.GetKafkaWriter(p.KafkaBroker, topic)
	kafkaMessages := make([]kafka.Message, 0)

	for i := 0; i < len(data); i++ {
		kafkaMessages = append(kafkaMessages, kafka.Message{
			Topic: data[i].Topic,
			Key:   data[i].Key,
			Value: data[i].Value,
		})
	}

	return kafkaWriter.WriteMessages(ctx, kafkaMessages...)
}
