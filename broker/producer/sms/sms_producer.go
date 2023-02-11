package sms

import (
	"context"
	"github.com/segmentio/kafka-go"
	"notification-service/broker/domain"
	"notification-service/config"
	"notification-service/utils"
)

type SMSProducerConfig struct {
	KafkaBroker []string
}

type SMSProducer interface {
	WriteMessage(ctx context.Context, data ...domain.QueueMessage) error
}

func NewSMSProducer(kafkaBroker []string) *SMSProducerConfig {
	return &SMSProducerConfig{
		KafkaBroker: kafkaBroker,
	}
}

func (p SMSProducerConfig) WriteMessage(ctx context.Context, data ...domain.QueueMessage) error {
	topic := config.AppConfig.KafkaSMSTopic
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
