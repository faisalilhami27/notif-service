package messageLog

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MessageLog struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	MessageID      string             `bson:"message_id" json:"message_id"`
	ReceiverNumber string             `bson:"receiver_number" json:"receiver_number"`
	ErrorMessage   string             `bson:"error_message" json:"error_message"`
	MessageRaw     interface{}        `bson:"message_raw" json:"message_raw"`
	Provider       string             `bson:"provider" json:"provider"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
}
