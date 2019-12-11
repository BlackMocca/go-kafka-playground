package event_handler

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
	"gitlab.com/km/go-kafka-playground/service/kafka"
)

type KafkaEventHandler struct {
	kafkaEventUs kafka.KafkaEventUsecase
}

func NewKafkaEventHandler(ev kafka.KafkaEventUsecase) *KafkaEventHandler {
	return &KafkaEventHandler{
		kafkaEventUs: ev,
	}
}

func (k KafkaEventHandler) Setup(_ sarama.ConsumerGroupSession) error { return nil }

func (k KafkaEventHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (k KafkaEventHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d value:%q \n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		log.Println("test test")
		sess.MarkMessage(msg, "")
	}
	return nil
}
