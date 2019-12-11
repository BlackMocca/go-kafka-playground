package kafka

import "github.com/Shopify/sarama"

type KafkaConsumerRepository interface {
}

type KafkaProducerRepository interface {
	PrepareMessage(topic, message string) *sarama.ProducerMessage
	SendOneMessageWithSync(msg *sarama.ProducerMessage) (int32, int64, error) // patition, offset, error
}
