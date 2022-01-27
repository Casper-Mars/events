package order

import (
	"context"
	"events/test/api"
	"log"
)

type orderCreateEventSub struct {
}

func NewOrderCreateEventSub() api.OrderCreateEventSubscriber {
	return &orderCreateEventSub{}
}

func (o *orderCreateEventSub) Receive(ctx context.Context, in *api.OrderCreateEvent) error {

	log.Printf("Receive OrderCreateEvent: %v", in)

	return nil
}
