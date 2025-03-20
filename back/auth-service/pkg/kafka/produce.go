package kafka

import (
	"auth/pkg/configs"
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaProducerService struct {
	Writer *kafka.Writer
	Topic  string
}

func NewKafkaProducerService(config *configs.Config) (*KafkaProducerService, error) {

	w := &kafka.Writer{
		Addr:     kafka.TCP(config.KafkaBootstrapServers),
		Topic:    config.KafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}

	return &KafkaProducerService{
		Writer: w,
		Topic:  config.KafkaTopic,
	}, nil
}

func (p *KafkaProducerService) SendMessage(key, value string) error {

	msg := &kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
	}

	err := p.Writer.WriteMessages(context.Background(), *msg)
	if err != nil {
		log.Printf("kafka producer: send message error: %v", err)
		return err
	}

	log.Printf("Key: %s, Value: %s", key, string(msg.Value))
	return nil
}

func (p *KafkaProducerService) Close() error {
	if err := p.Writer.Close(); err != nil {
		log.Printf("Error closing writer: %v", err)
		return err
	}
	return nil
}
