package event_handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Shopify/sarama"
	"gitlab.com/km/go-kafka-playground/config"
	"gitlab.com/km/go-kafka-playground/service/kafka"
)

type KafkaEventHandler struct {
	kafkaConsumerGroup sarama.ConsumerGroup
	kafkaEventUs       kafka.KafkaEventUsecase
}

func NewKafkaEventHandler(kcg sarama.ConsumerGroup, ev kafka.KafkaEventUsecase) *KafkaEventHandler {
	return &KafkaEventHandler{
		kafkaConsumerGroup: kcg,
		kafkaEventUs:       ev,
	}
}

func (k KafkaEventHandler) getActionFromTopic(topic string) (int, int) {
	var model, action int
	topicData := strings.Split(topic, config.TOPIC_DETERMINE)
	model, action = config.MODEL(topicData[0]), config.ACTION(topicData[1])
	return model, action
}

func (k KafkaEventHandler) Setup(_ sarama.ConsumerGroupSession) error { return nil }

func (k KafkaEventHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (k KafkaEventHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d value:%q \n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		model, action := k.getActionFromTopic(msg.Topic)

		switch model {
		case config.MODELUSER:
			switch action {
			case config.ACTIONCREATE:
				k.eventUserCreateHandler(sess, msg)
				sess.MarkMessage(msg, "")
			case config.ACTIONUPDATE:
				k.eventUserUpdateHandler(sess, msg)
				sess.MarkMessage(msg, "")
			case config.ACTIONDELETE:
				k.eventUserDeleteHandler(sess, msg)
				sess.MarkMessage(msg, "")
			}
		default:
			/** ไม่เข้าข่ายไหนเลย */
			sess.MarkMessage(msg, "")
		}
	}
	return nil
}

func (k KafkaEventHandler) eventUserCreateHandler(sess sarama.ConsumerGroupSession, message *sarama.ConsumerMessage) error {
	var val = string(message.Value)
	var id, _ = strconv.Atoi(val)

	if err := k.kafkaEventUs.CreateUserIntoMongo(id); err != nil {
		return err
	}

	return nil
}

func (k KafkaEventHandler) eventUserUpdateHandler(sess sarama.ConsumerGroupSession, message *sarama.ConsumerMessage) error {
	return nil
}

func (k KafkaEventHandler) eventUserDeleteHandler(sess sarama.ConsumerGroupSession, message *sarama.ConsumerMessage) error {
	return nil
}
