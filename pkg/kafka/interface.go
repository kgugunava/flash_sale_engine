package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaMessageInterface interface {
	GetMessage() kafka.Message
}

type KafkaProducerInterface interface {
	WriteMessages(ctx context.Context, msgs ...KafkaMessageInterface) error 
}

type KafkaConsumerInterface interface {
	ReadMessage(ctx context.Context) KafkaMessageInterface
}