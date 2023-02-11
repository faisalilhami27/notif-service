package sms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
	"notification-service/config"
	messageLogDT "notification-service/interfaces/dataTransfer/v1/messageLog"
	smsDT "notification-service/interfaces/dataTransfer/v1/sms"
	"notification-service/src/usecases/v1/messageLog"
	"notification-service/utils"
)

type SMSConsumerConfig struct {
	KafkaBroker []string
	useCase     messageLog.MessageLogUseCase
}

type SMSResponse struct {
	Data    map[string]interface{} `json:"data"`
	Code    int                    `json:"code"`
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
}

type SMSConsumer interface {
	ReadMessage(ctx context.Context)
}

func NewSMSConsumer(kafkaBroker []string, useCase messageLog.MessageLogUseCase) *SMSConsumerConfig {
	return &SMSConsumerConfig{
		KafkaBroker: kafkaBroker,
		useCase:     useCase,
	}
}

func (w SMSConsumerConfig) ReadMessage(ctx context.Context) {
	topic := config.AppConfig.KafkaSMSTopic
	groupID := config.AppConfig.KafkaSMSGroupId
	r := utils.GetKafkaReader(w.KafkaBroker, topic, groupID)

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			err := utils.GetErrorLog(err)
			if err != nil {
				return
			}
			break
		}

		message, err := utils.ParseQueueMessage(m.Value)
		if err != nil {
			err := utils.GetErrorLog(err)
			if err != nil {
				return
			}
		}
		byteData, err := json.Marshal(message.Data)
		if err != nil {
			err := utils.GetErrorLog(err)
			if err != nil {
				return
			}
			break
		}
		err = json.Unmarshal(byteData, &message.Data)
		if err != nil {
			err := utils.GetErrorLog(err)
			if err != nil {
				return
			}
		}
		body := smsDT.SMSBody{
			Message:         message.Data.(map[string]interface{})["message"].(string),
			RecipientNumber: message.Data.(map[string]interface{})["recipient_number"].(string),
		}
		err = sendMessageToTwilio(message.ID, body, w.useCase)
	}

	err := r.Close()
	if err != nil {
		err := utils.GetErrorLog(err)
		if err != nil {
			return
		}
	}
}

func sendMessageToTwilio(uuid string, body smsDT.SMSBody, useCase messageLog.MessageLogUseCase) error {
	twilioAuthToken := config.AppConfig.TwilioAuthToken
	twilioAccountId := config.AppConfig.TwilioAccountId
	twilioSMSSenderPhone := config.AppConfig.TwilioSMSSenderPhone
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: twilioAccountId,
		Password: twilioAuthToken,
	})

	if !utils.IsValidPhoneNumber(body.RecipientNumber) {
		err := utils.GetErrorLog(errors.New("invalid phone number"))
		if err != nil {
			return err
		}
	}

	params := &api.CreateMessageParams{}
	params.SetTo("+" + body.RecipientNumber)
	params.SetFrom("+" + twilioSMSSenderPhone)
	params.SetBody(body.Message)
	_, err := client.Api.CreateMessage(params)
	var log messageLogDT.MessageLogDT
	log = messageLogDT.MessageLogDT{
		MessageID:      uuid,
		Provider:       "sms",
		MessageRaw:     body,
		ReceiverNumber: body.RecipientNumber,
	}
	if err != nil {
		log.ErrorMessage = err.Error()
		utils.GetErrorLog(err)
	} else {
		log.ErrorMessage = ""
		fmt.Println("Message sent successfully!")
	}

	_, err = useCase.Create(log)
	if err != nil {
		utils.GetErrorLog(err)
		return err
	}
	return nil
}
