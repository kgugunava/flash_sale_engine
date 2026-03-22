package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
	cfg KafkaProducerConfig
}

func NewKafkaProducer(cfg KafkaProducerConfig) KafkaProducerInterface {
	return &KafkaProducer{
		writer: &kafka.Writer{
			Addr: kafka.TCP(cfg.Brokers...),
			Topic: cfg.Topic,
		},
		cfg: cfg,
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

func (p *KafkaProducer) Close() {
	if err := p.writer.Close(); err != nil {
		log.Fatal("failed to close producer writer: ", err)
	}
}