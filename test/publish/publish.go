package main

import (
	"context"
	"events/events"
	"events/test/api"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

var size = 100

func main() {
	sender := createSender()
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
		//channelEventPub.SendEvent(context.Background(), &channel.EnterEvent{
		//	Name: fmt.Sprintf("%d", i),
		//	Uid:  rand.Uint32(),
		//})
		log.Println("sent event")
		time.Sleep(time.Second)
	}
}

func createSender() events.Sender {
	//config := sarama.NewConfig()
	//config.Producer.Return.Successes = true
	//cli, err := sarama.NewClient([]string{"localhost:9093"}, config)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//sender, err := events.NewKafkaSender(cli)
	//if err != nil {
	//	log.Fatal(err)
	//}
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	sender := events.NewRabbitMQSender(connection)
	return sender
}
