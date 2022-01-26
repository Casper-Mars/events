package main

import (
	"events/api"
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"log"
	"time"
)

func main() {
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9093"}, sarama.NewConfig())
	if err != nil {
		panic(err)
	}
	size := 100
	go func() {
		for {
			err = <-producer.Errors()
			if err != nil {
				log.Printf("producer error: %v", err)
			}
		}
	}()
	for i := 0; i < size; i++ {
		event := api.Event{
			OrderId: int64(i),
			UserId:  int64(i),
		}
		marshal, err := proto.Marshal(&event)
		if err != nil {
			log.Printf("marshal error: %v", err)
			continue
		}
		producer.Input() <- &sarama.ProducerMessage{Topic: "order", Value: sarama.ByteEncoder(marshal)}

		log.Printf("send msg")
		time.Sleep(time.Second)
	}
	producer.Close()
}

//func test() {
//	cg, err := sarama.NewConsumerGroup([]string{"localhost:9093"}, "test_group", sarama.NewConfig())
//	if err != nil {
//		panic(err)
//	}
//
//}
