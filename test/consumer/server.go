package main

import (
	"events/api"
	"events/events"
	"github.com/Shopify/sarama"
)

func NewSubscriber(orderSub1 api.ReceiverServer, orderSub2 api.ReceiverServer) events.Subscriber {
	cli, err := sarama.NewClient([]string{"localhost:9093"}, sarama.NewConfig())
	if err != nil {
		panic(err)
	}
	kafkaServer := events.NewKafkaSubscriber(cli)
	api.RegisterHandler(kafkaServer, events.SubRequest{
		Topic: "order",
	}, orderSub1)
	api.RegisterHandler(kafkaServer, events.SubRequest{
		Topic: "order",
	}, orderSub2)
	return kafkaServer
}
