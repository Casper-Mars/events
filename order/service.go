package order

import (
	context "context"
	"events/api"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type OrderEvent struct {
	api.UnimplementedReceiverServer
}

func (o *OrderEvent) ReceiveEvent(ctx context.Context, event *api.Event) (*emptypb.Empty, error) {
	log.Printf("Received event: %v", event)
	return &emptypb.Empty{}, nil
}
