package domain

import "encoding/json"

type KafkaMessage struct {
	ID   string      `json:"id"`
	Data interface{} `json:"data"`
}

type QueueMessage struct {
	Topic string
	Key   []byte
	Value []byte
}

func (msg *KafkaMessage) MarshalBinary() (data []byte, err error) {
	return json.Marshal(msg)
}

func (msg *KafkaMessage) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, msg)
}
