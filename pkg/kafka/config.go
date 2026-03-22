package kafka

type KafkaProducerConfig struct {
	Brokers []string
	Topic string
}

type KafkaConsumerConfig struct {
	Brokers []string
	Topic string
	GroupID string
}