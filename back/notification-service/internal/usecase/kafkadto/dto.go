package kafkadto

type BaseEvent struct {
	Event string `json:"event"`
}

type UserCreatedEvent struct {
	Event            string `json:"event"`
	Email            string `json:"email"`
	VerificationCode string `json:"verification_code"`
}
