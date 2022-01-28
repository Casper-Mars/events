package events

import (
	"context"
	"github.com/streadway/amqp"
)

type rabbitMQReceiver struct {
	amqpChan  *amqp.Channel
	msgChan   <-chan amqp.Delivery
	queueName string
}

func (r *rabbitMQReceiver) Receive(ctx context.Context) (Message, error) {
	msg := <-r.msgChan

	return Message{
		Data:  msg.Body,
		Topic: r.queueName,
		Metadata: map[string]interface{}{
			"DeliveryTag": msg.DeliveryTag,
		},
	}, nil
}

func (r *rabbitMQReceiver) Ack(msg Message) error {
	tag := msg.Metadata["DeliveryTag"].(uint64)
	return r.amqpChan.Ack(tag, false)
}

func (r *rabbitMQReceiver) Nack(msg Message) error {
	tag := msg.Metadata["DeliveryTag"].(uint64)
	return r.amqpChan.Nack(tag, false, false)
}

func (r *rabbitMQReceiver) Close() error {
	return r.amqpChan.Close()
}

type rabbitMQReceiverBuilder struct {
	conn *amqp.Connection
}

func NewRabbitMQReceiverBuilder(conn *amqp.Connection) ReceiverBuilder {
	return &rabbitMQReceiverBuilder{
		conn: conn,
	}
}

func (r *rabbitMQReceiverBuilder) Build(subReq SubRequest) (Receiver, error) {
	channel, err := r.conn.Channel()
	if err != nil {
		return nil, err
	}
	consume, err := channel.Consume(subReq.Topic, "", false, false, false, false, nil)
	return &rabbitMQReceiver{
		amqpChan:  channel,
		msgChan:   consume,
		queueName: subReq.Topic,
	}, nil
}

type rabbitMQSender struct {
	amqpChan *amqp.Channel
}

func NewRabbitMQSender(conn *amqp.Connection) Sender {
	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return &rabbitMQSender{
		amqpChan: channel,
	}
}

func (r *rabbitMQSender) Send(ctx context.Context, message Message) error {
	return r.amqpChan.Publish("", message.Topic, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        message.Data,
	})
}

func (r *rabbitMQSender) Close() error {
	return r.amqpChan.Close()
}
