package main

import (
	"context"
	"events/order"
	"log"
)

func main() {
	subscriber := NewSubscriber(&order.EventSub1{}, &order.EventSub2{})
	subscriber.Start(context.Background())
	log.Printf("Kafka server started")
	select {}
}
