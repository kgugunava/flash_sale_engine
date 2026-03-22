package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(brokersAdrs []string, topic string) KafkaProducerInterface {
	return &KafkaProducer{
		writer: &kafka.Writer{
			Addr: kafka.TCP(brokersAdrs...),
			Topic: topic,
		},
	}
}

func (p *KafkaProducer) WriteMessages(ctx context.Context, msgs ...KafkaMessageInterface) error {
	for _, msg := range(msgs) {
		err := p.writer.WriteMessages(ctx, msg.GetMessage())
		if err != nil {
			log.Fatal("failed to write message: %s", msg.GetMessage().Key, err)
		}
	}

	return nil
}