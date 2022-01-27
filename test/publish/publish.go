package main

import (
	"context"
	"events/events"
	"events/test/api"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

var size = 100

func main() {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	cli, err := sarama.NewClient([]string{"localhost:9093"}, config)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err2 := cli.Close()
		if err2 != nil {
			log.Printf("Error closing client: %v", err2)
		}
	}()
	sender, err := events.NewKafkaSender(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err2 := sender.Close()
		if err2 != nil {
			log.Printf("Error closing sender: %v", err2)
		}
	}()
	buyEventPublisher := api.NewOrderCreateEventPublisher(events.PublishMetadata{
		Topic: "buy",
	}, sender)

	for i := 1; i <= size; i++ {
		err := buyEventPublisher.SendEvent(context.Background(), &api.OrderCreateEvent{
			Id:   int64(i),
			Name: fmt.Sprintf("%d", i),
		})
		if err != nil {
			log.Printf("error: %v", err)
		}
		log.Println("sent event")
		time.Sleep(time.Second)
	}
}
