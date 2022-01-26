package events

import (
	"context"
	"github.com/Shopify/sarama"
	"log"
)

type KafkaServer struct {
	consumer   sarama.Consumer
	handlerMap map[string]Handler
	done       chan struct{}
}

func NewKafkaServer(consumer sarama.Consumer) *KafkaServer {
	return &KafkaServer{
		consumer:   consumer,
		done:       make(chan struct{}),
		handlerMap: make(map[string]Handler),
	}
}

func (k *KafkaServer) AddHandler(event string, handler Handler) {
	k.handlerMap[event] = handler
}

func (k *KafkaServer) Start(ctx context.Context) error {
	for topic, handler := range k.handlerMap {
		partitionConsumer, err := k.consumer.ConsumePartition(topic, 0, 0)
		if err != nil {
			return err
		}
		tmpTopic := topic
		go func(handler Handler) {
			log.Printf("Started consumer for topic %s", tmpTopic)
			defer func() {
				partitionConsumer.AsyncClose()
			}()
			for {
				select {
				case msg := <-partitionConsumer.Messages():
					handler.Handle(context.Background(), msg.Value)
				case <-k.done:
					return
				}
			}
		}(handler)
	}
	return nil
}

func (k *KafkaServer) Stop(ctx context.Context) error {
	close(k.done)
	k.consumer.Close()
	return nil
}
