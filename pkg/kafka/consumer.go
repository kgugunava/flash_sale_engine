package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	reader *kafka.Reader
	cfg KafkaConsumerConfig
}

func NewKafkaConsumer(cfg KafkaConsumerConfig) KafkaConsumerInterface {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: cfg.Brokers,
		Topic: cfg.Topic,
		GroupID: cfg.GroupID,
	})
	return &KafkaConsumer{
		reader: reader,
		cfg: cfg,
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

func (c *KafkaConsumer) Close() {
	if err := c.reader.Close(); err != nil {
		log.Fatal("failed to close producer writer: ", err)
	}
}