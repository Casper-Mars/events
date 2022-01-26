package events

import (
	"context"
	"github.com/Shopify/sarama"
	"log"
)

type KafkaSubscriber struct {
	cli        sarama.Client
	handlerMap map[string]Handler
	done       chan struct{}
}

func NewKafkaSubscriber(cli sarama.Client) Subscriber {
	return &KafkaSubscriber{
		cli:        cli,
		handlerMap: make(map[string]Handler),
		done:       make(chan struct{}),
	}
}

func (k KafkaSubscriber) Start(ctx context.Context) error {
	return nil
}

func (k KafkaSubscriber) Stop(ctx context.Context) error {
	close(k.done)
	return k.cli.Close()
}

func (k KafkaSubscriber) Subscribe(subReq SubRequest, handler Handler) error {
	consumer, err := sarama.NewConsumerFromClient(k.cli)
	partitionConsumer, err := consumer.ConsumePartition(subReq.Topic, 0, sarama.OffsetNewest)
	if err != nil {
		return err
	}
	go func(handler Handler) {
		log.Printf("Started consumer for topic %s", subReq.Topic)
		defer func() {
			partitionConsumer.AsyncClose()
		}()
		for {
			select {
			case msg := <-partitionConsumer.Messages():
				handler.Handle(context.Background(), Message{
					Data:  msg.Value,
					Topic: msg.Topic,
				})
			case <-k.done:
				return
			}
		}
	}(handler)
	return nil
}

type KafkaReceiver struct {
	consumer         sarama.Consumer
	consumePartition sarama.PartitionConsumer
	handlerMap       map[string]Handler
	done             chan struct{}
}

func NewKafkaReceiver(cli sarama.Client, request SubRequest) (Receiver, error) {
	c, err := sarama.NewConsumerFromClient(cli)
	if err != nil {
		return nil, err
	}
	consumePartition, err := c.ConsumePartition(request.Topic, 0, sarama.OffsetNewest)
	if err != nil {
		return nil, err
	}

	return &KafkaReceiver{
		consumer:         c,
		consumePartition: consumePartition,
		handlerMap:       make(map[string]Handler),
		done:             make(chan struct{}),
	}, nil
}

func (k *KafkaReceiver) Receive(ctx context.Context) (Message, error) {
	select {
	case <-ctx.Done():
		return Message{}, nil
	case msg := <-k.consumePartition.Messages():
		return Message{
			Topic: msg.Topic,
			Data:  msg.Value,
		}, nil
	}
}

func (k *KafkaReceiver) Ack(msg Message) error {
	//TODO implement me
	panic("implement me")
}

func (k *KafkaReceiver) Nack(msg Message) error {
	//TODO implement me
	panic("implement me")
}

func (k *KafkaReceiver) Close() error {
	k.consumePartition.AsyncClose()
	return k.consumer.Close()
}

type KafkaReceiverBuilder struct {
	cli sarama.Client
}

func NewKafkaReceiverBuilder(cli sarama.Client) ReceiverBuilder {
	return &KafkaReceiverBuilder{
		cli: cli,
	}
}

func (k *KafkaReceiverBuilder) Build(subReq SubRequest) (Receiver, error) {
	return NewKafkaReceiver(k.cli, subReq)
}
