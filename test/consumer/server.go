package main

import (
	"events/api"
	"events/events"
	"github.com/Shopify/sarama"
)

func NewSubscriber(orderSub1 api.ReceiverServer, orderSub2 api.ReceiverServer, buySub api.BuyEventReceiverServer) events.Subscriber {
	cli, err := sarama.NewClient([]string{"localhost:9093"}, sarama.NewConfig())
	if err != nil {
		panic(err)
	}
	server := events.NewServer(events.WithReceiverBuilder(events.NewKafkaReceiverBuilder(cli)))
	api.RegisterOrderHandler(server, events.SubRequest{
		Topic: "order",
	}, orderSub1)
	api.RegisterOrderHandler(server, events.SubRequest{
		Topic: "order",
	}, orderSub2)
	api.RegisterBuyEventHandler(server, events.SubRequest{
		Topic: "buy",
	}, buySub)
	return server
}
