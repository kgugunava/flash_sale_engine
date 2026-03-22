package kafka

import (
	"github.com/segmentio/kafka-go"
)

type KafkaMessage struct {
	message kafka.Message
}

func (m *KafkaMessage) GetMessage() kafka.Message {
	return m.message
}