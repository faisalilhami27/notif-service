package whatsapp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
	"notification-service/config"
	messageLogDT "notification-service/interfaces/dataTransfer/v1/messageLog"
	"notification-service/interfaces/dataTransfer/v1/whatsapp"
	"notification-service/src/usecases/v1/messageLog"
	"notification-service/utils"
)

type WhatsappConsumerConfig struct {
	KafkaBroker []string
	useCase     messageLog.MessageLogUseCase
}

type WhatsappConsumer interface {
	ReadMessage(ctx context.Context)
}

func NewWhatsappConsumer(kafkaBroker []string, useCase messageLog.MessageLogUseCase) *WhatsappConsumerConfig {
	return &WhatsappConsumerConfig{
		KafkaBroker: kafkaBroker,
		useCase:     useCase,
	}
}

func (w WhatsappConsumerConfig) ReadMessage(ctx context.Context) {
	topic := config.AppConfig.KafkaWATopic
	groupID := config.AppConfig.KafkaWAGroupId
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
			break
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

		body := whatsapp.WhatsappBody{
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

func sendMessageToTwilio(uuid string, body whatsapp.WhatsappBody, useCase messageLog.MessageLogUseCase) error {
	twilioAuthToken := config.AppConfig.TwilioAuthToken
	twilioAccountId := config.AppConfig.TwilioAccountId
	twilioWaSenderPhone := config.AppConfig.TwilioWaSenderPhone
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
	params.SetTo("whatsapp:" + body.RecipientNumber)
	params.SetFrom("whatsapp:" + twilioWaSenderPhone)
	params.SetBody(body.Message)
	_, err := client.Api.CreateMessage(params)
	var log messageLogDT.MessageLogDT
	log = messageLogDT.MessageLogDT{
		MessageID:      uuid,
		Provider:       "whatsapp",
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
		err := utils.GetErrorLog(err)
		if err != nil {
			return err
		}
	}
	return nil
}
