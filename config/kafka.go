package config

import (
	"log"
	"os"
	"strings"

	"github.com/Shopify/sarama"
)

func initConfig() *sarama.Config {
	return sarama.NewConfig()
}

func getBrokerURLFromENV() []string {
	determine := ","
	kafkas := os.Getenv("KAFKA_URL")
	if strings.Index(kafkas, determine) != -1 {
		return strings.Split(kafkas, determine)
	}

	return []string{kafkas}
}

func NewBroker(url string) {
	broker := sarama.NewBroker(url)
	err := broker.Open(nil)
	if err != nil {
		panic(err)
	}
	connected, _ := broker.Connected()
	if !connected {
		log.Fatal("cant connect broker at ", url)
	}
}

func NewKafkaClient() sarama.Client {
	brokers := getBrokerURLFromENV()
	config := initConfig()
	client, err := sarama.NewClient(brokers, config)
	if err != nil {
		log.Fatal("error at Kafka client ", err)
	}
	return client
}
