package kafka

type KafkaUsecase interface {
	SendMessage(topic, message string) (int32, int64, error)
}

type KafkaEventUsecase interface {
}
