package whatsapp

import "encoding/json"

type SMSBody struct {
	Message         string `bson:"message" json:"message"`
	RecipientNumber string `bson:"recipient_number" json:"recipient_number"`
}

func (msg *SMSBody) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, msg)
}

func (msg *SMSBody) MarshalBinary() (data []byte, err error) {
	return json.Marshal(msg)
}
