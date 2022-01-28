package events

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"math/rand"
)

// kafka PoC

type KafkaReceiver struct {
	consumer sarama.ConsumerGroup
	msgCh    chan Message
	cancel   context.CancelFunc
}

type kafkaConsumerGroupHandler struct {
	msgCh chan Message
	ctx   context.Context
}

func (k *kafkaConsumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (k *kafkaConsumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (k *kafkaConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	log.Print("ready to receive event")
	for {
		select {
		case <-k.ctx.Done():
			return nil
		case msg := <-claim.Messages():
			if msg != nil {
				k.msgCh <- Message{
					Topic: msg.Topic,
					Data:  msg.Value,
				}
			}
		}
	}
}

func NewKafkaReceiver(cli sarama.Client, request SubRequest) (Receiver, error) {
	groupConsumer, err := sarama.NewConsumerGroupFromClient(fmt.Sprintf("%d", rand.Uint64()), cli)
	if err != nil {
		return nil, err
	}
	msgCh := make(chan Message)
	ctx, cancelFunc := context.WithCancel(context.Background())
	h := &kafkaConsumerGroupHandler{
		msgCh: msgCh,
		ctx:   ctx,
	}
	go func() {
		err = groupConsumer.Consume(context.Background(), []string{request.Topic}, h)
		if err != nil {
			log.Printf("error consuming: %v", err)
		}
	}()
	k := &KafkaReceiver{
		msgCh:    msgCh,
		consumer: groupConsumer,
		cancel:   cancelFunc,
	}
	return k, nil
}

func (k *KafkaReceiver) Receive(ctx context.Context) (Message, error) {
	select {
	case <-ctx.Done():
		return Message{}, nil
	case msg := <-k.msgCh:
		return msg, nil
	}
}

func (k *KafkaReceiver) Ack(msg Message) error {
	return nil
}

func (k *KafkaReceiver) Nack(msg Message) error {
	return nil
}

func (k *KafkaReceiver) Close() error {
	k.cancel()
	close(k.msgCh)
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
