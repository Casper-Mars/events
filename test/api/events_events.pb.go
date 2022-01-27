// Code generated by protoc-gen-go-events. DO NOT EDIT.

package api

import (
	context "context"
	events "events/events"
	proto "github.com/golang/protobuf/proto"
)

// OrderCreateEventPublisher is an interface to publish OrderCreateEvent event
type OrderCreateEventPublisher interface {
	SendEvent(ctx context.Context, in *OrderCreateEvent) error
}

// NewOrderCreateEventPublisher create a publisher for OrderCreateEvent event
func NewOrderCreateEventPublisher(metadata events.PublishMetadata, sender events.Sender) OrderCreateEventPublisher {
	return &orderCreateEventPublisher{
		sender:   sender,
		metadata: metadata,
	}
}

// OrderCreateEventSubscriber is an interface to subscribe to OrderCreateEvent event
type OrderCreateEventSubscriber interface {
	Receive(ctx context.Context, in *OrderCreateEvent) error
}

// RegisterOrderCreateEventSubscriber register a subscriber
func RegisterOrderCreateEventSubscriber(server events.Subscriber, subReq events.SubRequest, srv OrderCreateEventSubscriber) error {
	handler := &orderCreateEventHandler{
		srv: srv,
	}
	return server.Subscribe(subReq, handler)
}

type orderCreateEventPublisher struct {
	sender   events.Sender
	metadata events.PublishMetadata
}

func (b *orderCreateEventPublisher) SendEvent(ctx context.Context, in *OrderCreateEvent) error {
	data, err := proto.Marshal(in)
	if err != nil {
		return err
	}
	return b.sender.Send(ctx, events.Message{
		Data:  data,
		Topic: b.metadata.Topic,
	})
}

type orderCreateEventHandler struct {
	srv OrderCreateEventSubscriber
}

func (o *orderCreateEventHandler) Handle(ctx context.Context, msg events.Message) error {
	event := &OrderCreateEvent{}
	err := proto.Unmarshal(msg.Data, event)
	if err != nil {
		return err
	}
	return o.srv.Receive(ctx, event)
}

// OrderUpdateEventPublisher is an interface to publish OrderUpdateEvent event
type OrderUpdateEventPublisher interface {
	SendEvent(ctx context.Context, in *OrderUpdateEvent) error
}

// NewOrderUpdateEventPublisher create a publisher for OrderUpdateEvent event
func NewOrderUpdateEventPublisher(metadata events.PublishMetadata, sender events.Sender) OrderUpdateEventPublisher {
	return &orderUpdateEventPublisher{
		sender:   sender,
		metadata: metadata,
	}
}

// OrderUpdateEventSubscriber is an interface to subscribe to OrderUpdateEvent event
type OrderUpdateEventSubscriber interface {
	Receive(ctx context.Context, in *OrderUpdateEvent) error
}

// RegisterOrderUpdateEventSubscriber register a subscriber
func RegisterOrderUpdateEventSubscriber(server events.Subscriber, subReq events.SubRequest, srv OrderUpdateEventSubscriber) error {
	handler := &orderUpdateEventHandler{
		srv: srv,
	}
	return server.Subscribe(subReq, handler)
}

type orderUpdateEventPublisher struct {
	sender   events.Sender
	metadata events.PublishMetadata
}

func (b *orderUpdateEventPublisher) SendEvent(ctx context.Context, in *OrderUpdateEvent) error {
	data, err := proto.Marshal(in)
	if err != nil {
		return err
	}
	return b.sender.Send(ctx, events.Message{
		Data:  data,
		Topic: b.metadata.Topic,
	})
}

type orderUpdateEventHandler struct {
	srv OrderUpdateEventSubscriber
}

func (o *orderUpdateEventHandler) Handle(ctx context.Context, msg events.Message) error {
	event := &OrderUpdateEvent{}
	err := proto.Unmarshal(msg.Data, event)
	if err != nil {
		return err
	}
	return o.srv.Receive(ctx, event)
}

// OrderDeleteEventPublisher is an interface to publish OrderDeleteEvent event
type OrderDeleteEventPublisher interface {
	SendEvent(ctx context.Context, in *OrderDeleteEvent) error
}

// NewOrderDeleteEventPublisher create a publisher for OrderDeleteEvent event
func NewOrderDeleteEventPublisher(metadata events.PublishMetadata, sender events.Sender) OrderDeleteEventPublisher {
	return &orderDeleteEventPublisher{
		sender:   sender,
		metadata: metadata,
	}
}

// OrderDeleteEventSubscriber is an interface to subscribe to OrderDeleteEvent event
type OrderDeleteEventSubscriber interface {
	Receive(ctx context.Context, in *OrderDeleteEvent) error
}

// RegisterOrderDeleteEventSubscriber register a subscriber
func RegisterOrderDeleteEventSubscriber(server events.Subscriber, subReq events.SubRequest, srv OrderDeleteEventSubscriber) error {
	handler := &orderDeleteEventHandler{
		srv: srv,
	}
	return server.Subscribe(subReq, handler)
}

type orderDeleteEventPublisher struct {
	sender   events.Sender
	metadata events.PublishMetadata
}

func (b *orderDeleteEventPublisher) SendEvent(ctx context.Context, in *OrderDeleteEvent) error {
	data, err := proto.Marshal(in)
	if err != nil {
		return err
	}
	return b.sender.Send(ctx, events.Message{
		Data:  data,
		Topic: b.metadata.Topic,
	})
}

type orderDeleteEventHandler struct {
	srv OrderDeleteEventSubscriber
}

func (o *orderDeleteEventHandler) Handle(ctx context.Context, msg events.Message) error {
	event := &OrderDeleteEvent{}
	err := proto.Unmarshal(msg.Data, event)
	if err != nil {
		return err
	}
	return o.srv.Receive(ctx, event)
}