package main

import (
	"context"
	"events/test/order"
	"log"
)

func main() {
	subscriber := NewSubscriber(order.NewOrderCreateEventSub())
	err := subscriber.Start(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Kafka server started")
	select {}
}
