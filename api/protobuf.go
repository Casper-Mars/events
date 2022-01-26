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

func (o *OrderHandler) Handle(ctx context.Context, msg events.Message) {
	event := &Event{}
	err := proto.Unmarshal(msg.Data, event)
	if err != nil {
		log.Printf("handle event error: %v", err)
		return
	}
	_, _ = o.srv.ReceiveEvent(ctx, event)
}

func RegisterHandler(server events.Subscriber, subReq events.SubRequest, srv ReceiverServer) {
	err := server.Subscribe(subReq, NewOrderHandler(srv))
	if err != nil {
		log.Fatalf("subscribe error: %v", err)
	}
}
