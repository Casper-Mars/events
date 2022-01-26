package api

import (
	"context"
	"events/events"
	"github.com/golang/protobuf/proto"
	"log"
)

type OrderHandler struct {
	srv ReceiverServer
}

func NewOrderHandler(srv ReceiverServer) events.Handler {
	return &OrderHandler{
		srv: srv,
	}
}

func (o *OrderHandler) Handle(ctx context.Context, data []byte) {
	event := &Event{}
	err := proto.Unmarshal(data, event)
	if err != nil {
		log.Printf("handle event error: %v", err)
		return
	}
	_, _ = o.srv.ReceiveEvent(ctx, event)
}

func RegisterHandler(server *events.KafkaServer, topic string, srv ReceiverServer) {
	server.AddHandler(topic, NewOrderHandler(srv))
}

func RegisterHandlerTmp(server *events.Subscriber, topic string, srv ReceiverServer) {
	server.AddHandler(topic, NewOrderHandler(srv))
}
