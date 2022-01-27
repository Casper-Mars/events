package user

import (
	"context"
	"events/test/api/user"
	"log"
)

type Events struct {
}

func NewEvents() *Events {
	return &Events{}
}

func (e *Events) Receive(ctx context.Context, in *user.LoginEvent) error {
	log.Printf("Received event: %v", in)
	return nil
}
