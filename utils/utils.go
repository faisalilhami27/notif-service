package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"log"
	"notification-service/broker/domain"
	"regexp"
	"time"
)

func LoadEnv() error {
	env := godotenv.Load()
	if env != nil {
		log.Fatal("Error loading .env file")
	}

	return env
}

func GetKafkaWriter(addresses []string, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(addresses...),
		Topic:        topic,
		WriteTimeout: 1 * time.Second,
	}
}

func ParseQueueMessage(msg []byte) (*domain.KafkaMessage, error) {
	qm := domain.KafkaMessage{}
	err := qm.UnmarshalBinary(msg)
	if err != nil {
		return nil, err
	}

	return &qm, nil
}

func GetKafkaReader(addresses []string, topic string, kafkaConsumerGroupId string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: addresses,
		Topic:   topic,
		GroupID: kafkaConsumerGroupId,
		MaxWait: 1 * time.Second,
	})
}

func IsValidPhoneNumber(phone string) bool {
	// This regex will handle (62812...) and (+62812...)
	// Turns out, if you use '+' in requestParameter, it will be changed into space ' ',
	// this regex handle that case as well
	phoneRegex := "^[\\+ ]{0,1}[1-9][0-9]{5,15}$"

	r, _ := regexp.MatchString(phoneRegex, phone)
	return r
}

func GetErrorLog(err error) error {
	fmt.Println(err)
	return err
}
