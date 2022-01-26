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

func (o *OrderHandler) Handle(ctx context.Context, msg events.Message) error {
	event := &Event{}
	err := proto.Unmarshal(msg.Data, event)
	if err != nil {
		log.Printf("handle event error: %v", err)
		return err
	}
	_, err = o.srv.ReceiveEvent(ctx, event)
	return err
}

func RegisterOrderHandler(server events.Subscriber, subReq events.SubRequest, srv ReceiverServer) {
	err := server.Subscribe(subReq, NewOrderHandler(srv))
	if err != nil {
		log.Fatalf("subscribe error: %v", err)
	}
}

type BuyEventHandler struct {
	srv BuyEventReceiverServer
}

func NewBuyEventHandler(srv BuyEventReceiverServer) events.Handler {
	return &BuyEventHandler{
		srv: srv,
	}
}

func (b *BuyEventHandler) Handle(ctx context.Context, msg events.Message) error {
	buyEvent := &BuyEvent{}
	err := proto.Unmarshal(msg.Data, buyEvent)
	if err != nil {
		log.Printf("handle buy event error: %v", err)
		return err
	}
	_, err = b.srv.ReceiveEvent(ctx, buyEvent)
	return err
}

func RegisterBuyEventHandler(server events.Subscriber, subReq events.SubRequest, srv BuyEventReceiverServer) {
	err := server.Subscribe(subReq, NewBuyEventHandler(srv))
	if err != nil {
		log.Fatalf("subscribe error: %v", err)
	}
}
