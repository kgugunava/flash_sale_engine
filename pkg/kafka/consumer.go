package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	reader *kafka.Reader
}

func NewKafkaConsumer(brokerAddrs []string, topic string, groupID string) KafkaConsumerInterface {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokerAddrs,
		Topic: topic,
		GroupID: groupID,
	})
	return &KafkaConsumer{
		reader: reader,
	}
}

func (c *KafkaConsumer) ReadMessage(ctx context.Context) KafkaMessageInterface {
	msg, err := c.reader.ReadMessage(ctx)
	if err != nil {
		log.Fatal("failed to read message from topic: %s", c.reader.Config().Topic, err)
	}
	return &KafkaMessage{
		message: msg,
	}
}