package config

import (
	"log"

	"github.com/Shopify/sarama"
)

type KafkaConsumerGroup struct {
	consumerGroup sarama.ConsumerGroup
}

func settingConsumerGroupConfig(config *sarama.Config) {
	config.Version = sarama.V2_3_0_0
	config.Consumer.Return.Errors = true
}

func NewKafkaConsumerGroupFromClient(group string, client sarama.Client) *KafkaConsumerGroup {
	settingConsumerGroupConfig(client.Config())

	consumerGroup, err := sarama.NewConsumerGroupFromClient(group, client)
	if err != nil {
		log.Fatal(err)
	}

	return &KafkaConsumerGroup{
		consumerGroup: consumerGroup,
	}
}

func (k KafkaConsumerGroup) GetConsumerGroup() sarama.ConsumerGroup {
	return k.consumerGroup
}
