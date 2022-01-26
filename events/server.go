package events

import (
	"context"
)

type subscriber struct {
}

func NewSubscriber() Subscriber {
	s := &subscriber{}
	return s
}

func (s *subscriber) Start(ctx context.Context) error {
	return nil
}

func (s *subscriber) Stop(ctx context.Context) error {
	return nil
}

func (s *subscriber) Subscribe(subReq SubRequest, handler Handler) error {
	return nil
}
