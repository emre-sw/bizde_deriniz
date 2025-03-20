package kafka

import (
	"encoding/json"
	"log"
	"notification/internal/usecase/kafkadto"
	"notification/pkg/email"
)

func ConsumeUserCreated(event []byte) error {
	var e kafkadto.UserCreatedEvent
	if err := json.Unmarshal(event, &e); err != nil {
		log.Printf("Error unmarshalling event: %v", err)
		return err
	}

	log.Printf("Processing user created event for email: %s", e.Email)

	if err := email.SendEmail(e.Email, e.VerificationCode); err != nil {
		log.Printf("Error sending email: %v", err)
		return err
	}

	log.Printf("Successfully processed user created event for email: %s", e.Email)
	return nil
}
