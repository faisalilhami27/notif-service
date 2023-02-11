package utils

import (
	"log"
)

func MongoErrorException(message error) bool {
	if message != nil {
		log.Printf("Caused error exceptions : %s", message.Error())
	}

	return message == nil
}

func MongoErrorDatabaseException(message error) bool {
	if message != nil {
		log.Printf("Caused error database exceptions : %s", message.Error())
	}

	return message == nil
}
