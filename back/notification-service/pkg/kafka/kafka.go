package kafka

import (
	"context"
	"encoding/json"
	"log"
	"notification/internal/usecase/kafkadto"
	"notification/pkg/configs"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumerService struct {
	reader *kafka.Reader
}

func NewKafkaConsumerService(config *configs.Config) *KafkaConsumerService {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{config.KafkaBootstrapServers},
		GroupID:     config.KafkaGroupID,
		Topic:       "user_created",
		StartOffset: kafka.FirstOffset,
		MinBytes:    10e3,
		MaxBytes:    10e6,
	})

	log.Printf("Created Kafka consumer for topic user_created with group ID %s", config.KafkaGroupID)

	return &KafkaConsumerService{
		reader: reader,
	}
}

func (k *KafkaConsumerService) Consume(topic string) {
	log.Printf("Starting to consume messages from topic: %s", topic)

	for {
		msg, err := k.reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		log.Printf("Received message from topic %s: %s", topic, string(msg.Value))

		var event kafkadto.BaseEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			continue
		}

		switch event.Event {
		case "user_created":
			if err := ConsumeUserCreated(msg.Value); err != nil {
				log.Printf("Error processing user_created event: %v", err)
			}
		default:
			log.Printf("Unknown event type: %s", event.Event)
		}
	}
}

func (k *KafkaConsumerService) Close() error {
	return k.reader.Close()
}
