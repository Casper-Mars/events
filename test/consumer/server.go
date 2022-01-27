package main

import (
	"events/events"
	"events/test/api"
	"github.com/Shopify/sarama"
)

func NewSubscriber(oce api.OrderCreateEventSubscriber) events.Subscriber {
	cli, err := sarama.NewClient([]string{"localhost:9093"}, sarama.NewConfig())
	if err != nil {
		panic(err)
	}
	server := events.NewServer(events.WithReceiverBuilder(events.NewKafkaReceiverBuilder(cli)))
	err = api.RegisterOrderCreateEventSubscriber(server, events.SubRequest{
		Topic: "buy",
	}, oce)
	if err != nil {
		panic(err)
	}
	return server
}
