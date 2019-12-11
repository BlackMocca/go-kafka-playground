package config

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
	helperReq "gitlab.com/km/go-kafka-playground/helper/request"
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

func (k KafkaConsumerGroup) Setup(_ sarama.ConsumerGroupSession) error { return nil }

func (k KafkaConsumerGroup) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (k KafkaConsumerGroup) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		sess.MarkMessage(msg, "")
		messageReceivedGroup(msg)
	}
	return nil
}

func messageReceivedGroup(message *sarama.ConsumerMessage) {
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
