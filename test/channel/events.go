package channel

import (
	"context"
	"events/test/api/channel"
	"log"
)

type Events struct {
}

func (e Events) Receive(ctx context.Context, in *channel.EnterEvent) error {
	log.Printf("Receive: %v", in)
	return nil
}
