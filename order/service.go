package order

import (
	context "context"
	"events/api"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type EventSub1 struct {
	api.UnimplementedReceiverServer
}

func (o *EventSub1) ReceiveEvent(ctx context.Context, event *api.Event) (*emptypb.Empty, error) {
	log.Printf("Sub1 Received event: %v", event)
	return &emptypb.Empty{}, nil
}

type EventSub2 struct {
	api.UnimplementedReceiverServer
}

func (o *EventSub2) ReceiveEvent(ctx context.Context, event *api.Event) (*emptypb.Empty, error) {
	log.Printf("Sub2 Received event: %v", event)
	return &emptypb.Empty{}, nil
}

type BuyEventSub struct {
	api.UnimplementedBuyEventReceiverServer
}

func (b *BuyEventSub) ReceiveEvent(ctx context.Context, event *api.BuyEvent) (*emptypb.Empty, error) {
	log.Printf("Buy Event Received event: %v", event)
	return &emptypb.Empty{}, nil
}
