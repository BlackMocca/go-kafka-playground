package config

import (
	"log"

	"github.com/Shopify/sarama"
	helperReq "gitlab.com/km/go-kafka-playground/helper/request"
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
			// for message := range pc.Messages() {
			// 	// messageReceived(message)
			// 	log.Println("one broker in message", message)
			// }
		}(pc)
	}
}

func messageReceived(message *sarama.ConsumerMessage) {
	log.Println("message receive ", (string(message.Value)))
	userId := string(message.Value)
	url := "http://127.0.0.1:3000/kafka/users/" + userId
	log.Println(url)

	data, err := helperReq.RequestGET(url, nil, nil)
	if err != nil {
		log.Fatal("req error ", err)
	}

	log.Println(data)
}
