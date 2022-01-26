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
