package config

import (
	"log"
	"os"
	"strings"

	"github.com/Shopify/sarama"
)

var (
	determine = ","
)

func initConfig() *sarama.Config {
	return sarama.NewConfig()
}

func GetProducerBrokers() []string {
	var producer []string
	producers := os.Getenv("PRODUCER_URL")

	if strings.Index(producers, determine) != -1 {
		producer = strings.Split(producers, determine)
	} else {
		producer = []string{producers}
	}
	return producer
}

func GetConsumerBrokers() []string {
	var consumer []string
	consumers := os.Getenv("CONSUMER_URL")

	if strings.Index(consumers, determine) != -1 {
		consumer = strings.Split(consumers, determine)
	} else {
		consumer = []string{consumers}
	}

	return consumer
}

func GetBrokersURLFromENV() ([]string, []string) {
	return GetProducerBrokers(), GetConsumerBrokers()
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

func NewKafkaClient() (sarama.Client, sarama.Client) {
	producers, consumers := GetBrokersURLFromENV()
	config := initConfig()

	log.Println("kafka producers url are", producers)
	clientProducer, err := sarama.NewClient(producers, config)
	if err != nil {
		log.Fatal("[Producers] ", err)
	}

	log.Println("kafka consumers url are", consumers)
	clientConsumer, err := sarama.NewClient(consumers, config)
	if err != nil {
		log.Fatal("[Consumers] ", err)
	}

	return clientProducer, clientConsumer
}
