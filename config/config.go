package config

import (
	"os"
	"strings"
)

var AppConfig Environment

type Environment struct {
	DBHost               string
	DBPort               string
	DBUsername           string
	DBPassword           string
	DBDatabase           string
	Port                 string
	TwilioAccountId      string
	TwilioAuthToken      string
	TwilioWaSenderPhone  string
	TwilioSMSSenderPhone string
	KafkaBrokerAddress   []string
	KafkaWATopic         string
	KafkaWAGroupId       string
	KafkaSMSTopic        string
	KafkaSMSGroupId      string
}

func InitApp() {
	AppConfig.DBDatabase = os.Getenv("DB_DATABASE")
	if AppConfig.DBDatabase == "" {
		panic("DB_DATABASE IS EMPTY, PLEASE CONFIGURE FIRST")
	}

	AppConfig.DBHost = os.Getenv("DB_HOST")
	if AppConfig.DBHost == "" {
		panic("DB_HOST IS EMPTY, PLEASE CONFIGURE FIRST")
	}

	AppConfig.DBPort = os.Getenv("DB_PORT")
	if AppConfig.DBPort == "" {
		panic("DB_PORT IS EMPTY, PLEASE CONFIGURE FIRST")
	}

	AppConfig.DBUsername = os.Getenv("DB_USERNAME")
	if AppConfig.DBUsername == "" {
		panic("DB_USERNAME IS EMPTY, PLEASE CONFIGURE FIRST")
	}

	AppConfig.DBPassword = os.Getenv("DB_PASSWORD")
	if AppConfig.DBPassword == "" {
		panic("DB_PASSWORD IS EMPTY, PLEASE CONFIGURE FIRST")
	}

	AppConfig.Port = os.Getenv("PORT")
	if AppConfig.Port == "" {
		panic("PORT IS EMPTY, PLEASE CONFIGURE FIRST")
	}

	AppConfig.KafkaBrokerAddress = strings.Split(os.Getenv("KAFKA_BROKER_ADDRESSES"), ",")
	if AppConfig.KafkaBrokerAddress[0] == "" {
		panic("KAFKA_BROKER_ADDRESS IS EMPTY, PLEASE CONFIGURE FIRST")
	}

	AppConfig.KafkaWATopic = os.Getenv("KAFKA_WA_TOPIC")
	if AppConfig.KafkaWATopic == "" {
		panic("KAFKA_WA_TOPIC IS EMPTY, PLEASE CONFIGURE FIRST")
	}

	AppConfig.KafkaSMSTopic = os.Getenv("KAFKA_SMS_TOPIC")
	if AppConfig.KafkaSMSTopic == "" {
		panic("KAFKA_SMS_TOPIC IS EMPTY, PLEASE CONFIGURE FIRST")
	}

	AppConfig.KafkaWAGroupId = os.Getenv("KAFKA_WA_GROUP_ID")
	if AppConfig.KafkaWAGroupId == "" {
		panic("KAFKA_WA_GROUP_ID IS EMPTY, PLEASE CONFIGURE FIRST")
	}

	AppConfig.KafkaSMSGroupId = os.Getenv("KAFKA_SMS_GROUP_ID")
	if AppConfig.KafkaSMSGroupId == "" {
		panic("KAFKA_SMS_GROUP_ID IS EMPTY, PLEASE CONFIGURE FIRST")
	}

	AppConfig.TwilioAccountId = os.Getenv("TWILIO_ACCOUNT_ID")
	if AppConfig.TwilioAccountId == "" {
		panic("TWILIO_ACCOUNT_ID IS EMPTY, PLEASE CONFIGURE FIRST")
	}

	AppConfig.TwilioAuthToken = os.Getenv("TWILIO_AUTH_TOKEN")
	if AppConfig.TwilioAuthToken == "" {
		panic("TWILIO_AUTH_TOKEN IS EMPTY, PLEASE CONFIGURE FIRST")
	}

	AppConfig.TwilioWaSenderPhone = os.Getenv("TWILIO_WA_SENDER_PHONE")
	if AppConfig.TwilioWaSenderPhone == "" {
		panic("TWILIO_WA_SENDER_PHONE IS EMPTY, PLEASE CONFIGURE FIRST")
	}

	AppConfig.TwilioSMSSenderPhone = os.Getenv("TWILIO_SMS_SENDER_PHONE")
	if AppConfig.TwilioSMSSenderPhone == "" {
		panic("TWILIO_SMS_SENDER_PHONE IS EMPTY, PLEASE CONFIGURE FIRST")
	}
}
