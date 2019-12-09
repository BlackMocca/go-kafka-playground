package config

import (
	"log"

	"github.com/Shopify/sarama"
)

type KafkaProducer struct {
	syncProducer  sarama.SyncProducer
	asyncProducer sarama.AsyncProducer
}

func settingSyncProducerConfig(config *sarama.Config) {
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
}

func NewKafkaSyncProducer(client sarama.Client) *KafkaProducer {
	settingSyncProducerConfig(client.Config())

	producerInf, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		log.Fatal(err)
	}

	return &KafkaProducer{
		syncProducer: producerInf,
	}
}

func NewKafkaASyncProducer(client sarama.Client) *KafkaProducer {
	producerInf, err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		log.Fatal(err)
	}

	return &KafkaProducer{
		asyncProducer: producerInf,
	}
}

func (k KafkaProducer) GetSyncProducer() sarama.SyncProducer {
	return k.syncProducer
}
