package main

import (
	"context"
	"events/api"
	"events/events"
	"events/order"
	"github.com/Shopify/sarama"
	"log"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9093"}, sarama.NewConfig())
	if err != nil {
		panic(err)
	}
	kafkaServer := events.NewKafkaServer(consumer)
	api.RegisterHandler(kafkaServer, "order", &order.OrderEvent{})
	err = kafkaServer.Start(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Kafka server started")
	select {}
}
