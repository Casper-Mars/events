package main

import (
	"context"
	"events/order"
	"log"
)

func main() {
	subscriber := NewSubscriber(&order.EventSub1{}, &order.EventSub2{}, &order.BuyEventSub{})
	subscriber.Start(context.Background())
	log.Printf("Kafka server started")
	select {}
}
