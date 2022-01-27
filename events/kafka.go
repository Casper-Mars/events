package events

import (
	"context"
	"github.com/Shopify/sarama"
)

// kafka PoC

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

type KafkaSender struct {
	producer sarama.SyncProducer
}

func NewKafkaSender(cli sarama.Client) (Sender, error) {
	producer, err := sarama.NewSyncProducerFromClient(cli)
	if err != nil {
		return nil, err
	}
	return &KafkaSender{
		producer: producer,
	}, nil
}

func (k *KafkaSender) Send(ctx context.Context, message Message) error {
	_, _, err := k.producer.SendMessage(&sarama.ProducerMessage{
		Topic: message.Topic,
		Value: sarama.ByteEncoder(message.Data),
	})
	return err
}

func (k *KafkaSender) Close() error {
	return k.producer.Close()
}
