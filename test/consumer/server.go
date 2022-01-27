package main

import (
	"events/events"
	"events/test/api"
	"events/test/api/channel"
	"events/test/api/user"
	"github.com/Shopify/sarama"
)

func NewSubscriber(oce api.OrderCreateEventSubscriber, lce user.LoginEventSubscriber, cce channel.EnterEventSubscriber) events.Subscriber {
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
	err = user.RegisterLoginEventSubscriber(server, events.SubRequest{
		Topic: "login",
	}, lce)
	if err != nil {
		panic(err)
	}
	err = channel.RegisterEnterEventSubscriber(server, events.SubRequest{
		Topic: "channel",
	}, cce)
	err = channel.RegisterEnterEventSubscriber(server, events.SubRequest{
		Topic: "channel",
	}, cce)

	return server
}
