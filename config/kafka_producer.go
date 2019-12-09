package config

import (
	"log"
	"time"

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

func settingAsyncProducerConfig(config *sarama.Config) {
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

func NewKafkaAsyncProducer(client sarama.Client) *KafkaProducer {
	settingAsyncProducerConfig(client.Config())

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

func (k KafkaProducer) GetAsyncProducer() sarama.AsyncProducer {
	return k.asyncProducer
}

func (k *KafkaProducer) SetSyncProducer(sync sarama.SyncProducer) {
	k.syncProducer = sync
}

func (k *KafkaProducer) SetAsyncProducer(async sarama.AsyncProducer) {
	k.asyncProducer = async
}

func (k *KafkaProducer) SetingAsyncProducer(client sarama.Client) {
	settingAsyncProducerConfig(client.Config())

	producerInf, err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		log.Fatal(err)
	}

	k.asyncProducer = producerInf
}

func (k KafkaProducer) PrepareMessage(topic, message string) *sarama.ProducerMessage {
	now := time.Now()
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
		Timestamp: now,
	}

	return msg
}
