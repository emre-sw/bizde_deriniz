package main

import (
	"log"
	"notification/pkg/configs"
	"notification/pkg/kafka"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	config, err := configs.LoadConfig(".env")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	consumer := kafka.NewKafkaConsumerService(config)

	var wg sync.WaitGroup

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		defer wg.Done()
		consumer.Consume("user_created")
	}()

	<-done
	log.Println("Shutting down...")

	wg.Wait()
}
