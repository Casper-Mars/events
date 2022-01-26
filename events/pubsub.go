package events

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport"
)

type Handler interface {
	Handle(ctx context.Context, msg Message)
}

type Message struct {
	Topic string
	Data  []byte
}

type SubRequest struct {
	Topic string
}

type Subscriber interface {
	transport.Server
	Subscribe(subReq SubRequest, handler Handler) error
}
