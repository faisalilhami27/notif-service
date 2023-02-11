package whatsapp

import "encoding/json"

type WhatsappBody struct {
	Message         string `bson:"message" json:"message"`
	RecipientNumber string `bson:"recipient_number" json:"recipient_number"`
}

func (msg *WhatsappBody) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, msg)
}

func (msg *WhatsappBody) MarshalBinary() (data []byte, err error) {
	return json.Marshal(msg)
}
