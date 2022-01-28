package main

import (
	"context"
	"events/test/channel"
	"events/test/order"
	"events/test/user"
	"log"
)

func main() {
	subscriber := NewSubscriber(order.NewOrderCreateEventSub(), &user.Events{}, &channel.Events{})
	err := subscriber.Start(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("server started")
	select {}
}
