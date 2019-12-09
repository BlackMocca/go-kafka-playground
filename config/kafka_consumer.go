package config

import (
	"log"

	"github.com/Shopify/sarama"
)

type KafkaConsumer struct {
	consumer sarama.Consumer
}

func initConsumerConfig() *sarama.Config {
	return sarama.NewConfig()
}

func NewKafkaConsumer(client sarama.Client) *KafkaConsumer {
	consumerInf, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Fatal(err)
	}

	return &KafkaConsumer{
		consumer: consumerInf,
	}
}

func (k KafkaConsumer) GetConsumer() sarama.Consumer {
	return k.consumer
}

func (k KafkaConsumer) Subscribe(topic string) {
	partitionList, err := k.consumer.Partitions(topic) //get all partitions on the given topic
	if err != nil {
		log.Fatal("Error retrieving partitionList ", err)
	}
	initialOffset := sarama.OffsetNewest //get offset for the oldest message on the topic

	for _, partition := range partitionList {
		/**
		create partition
		*/
		pc, err := k.consumer.ConsumePartition(topic, partition, initialOffset)
		if err != nil {
			continue
		}

		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				messageReceived(message)
			}
		}(pc)
	}
}

func messageReceived(message *sarama.ConsumerMessage) {
	log.Println("message receive ", (string(message.Value)))
}
