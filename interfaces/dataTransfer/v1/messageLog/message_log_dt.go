package messageLog

type MessageLogDT struct {
	MessageID      string      `bson:"message_id" json:"message_id"`
	ReceiverNumber string      `bson:"receiver_number" json:"receiver_number"`
	ErrorMessage   string      `bson:"error_message" json:"error_message"`
	MessageRaw     interface{} `bson:"message_raw" json:"message_raw"`
	Provider       string      `bson:"provider" json:"provider"`
}
