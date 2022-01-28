package main

import (
	"events/events"
	"events/test/api"
	"events/test/api/channel"
	"events/test/api/user"
	"github.com/Shopify/sarama"
)

func NewSubscriber(oce api.OrderCreateEventSubscriber, lce user.LoginEventSubscriber, cce channel.EnterEventSubscriber) events.Subscriber {
	server := events.NewServer(events.WithReceiverBuilder(createReceiverBuilder()))
	err := api.RegisterOrderCreateEventSubscriber(server, events.SubRequest{
		Topic: "buy",
	}, oce)
	if err != nil {
		panic(err)
	}
	//err = user.RegisterLoginEventSubscriber(server, events.SubRequest{
	//	Topic: "login",
	//}, lce)
	//if err != nil {
	//	panic(err)
	//}
	//err = channel.RegisterEnterEventSubscriber(server, events.SubRequest{
	//	Topic: "channel",
	//}, cce)
	//err = channel.RegisterEnterEventSubscriber(server, events.SubRequest{
	//	Topic: "channel",
	//}, cce)

	return server
}

func createReceiverBuilder() events.ReceiverBuilder {
	cli, err := sarama.NewClient([]string{"localhost:9093"}, sarama.NewConfig())
	if err != nil {
		panic(err)
	}
	builder := events.NewKafkaReceiverBuilder(cli)
	//dial, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	//if err != nil {
	//	panic(err)
	//}
	//builder := events.NewRabbitMQReceiverBuilder(dial)
	return builder
}
