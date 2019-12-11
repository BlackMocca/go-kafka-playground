package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Shopify/sarama"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.com/km/go-kafka-playground/config"
	_conf "gitlab.com/km/go-kafka-playground/config"

	myMiddL "gitlab.com/km/go-kafka-playground/middleware"
	_user_handler "gitlab.com/km/go-kafka-playground/service/user/http"
	_user_repository "gitlab.com/km/go-kafka-playground/service/user/repository"
	_user_usecase "gitlab.com/km/go-kafka-playground/service/user/usecase"

	_kafka_event_handler "gitlab.com/km/go-kafka-playground/service/kafka/event_handler"
	_kafka_handler "gitlab.com/km/go-kafka-playground/service/kafka/http"
	_kafka_repository "gitlab.com/km/go-kafka-playground/service/kafka/repository"
	_kafka_usecase "gitlab.com/km/go-kafka-playground/service/kafka/usecase"
)

var (
	Config              *_conf.Config
	KafkaProducerClient sarama.Client
	KafkaConsumerClient sarama.Client
	KafkaProducer       *_conf.KafkaProducer
	// KafkaConsumer       *_conf.KafkaConsumer
	KafkaConsumerGroup *_conf.KafkaConsumerGroup

	KafkaProducerAsync *_conf.KafkaProducer
)

func init() {
	Config = _conf.NewConfigWithService(_conf.NewPsqlConnection(), _conf.NewMongoSession())
	KafkaProducerClient, KafkaConsumerClient = _conf.NewKafkaClient()
	Config.KafkaProducerClient = KafkaProducerClient
	Config.KafkaConsumerClient = KafkaConsumerClient
}

func createTopic(topic string) {
	// KafkaConsumer.Subscribe(topic)
}

func main() {
	KafkaProducer = config.NewKafkaSyncProducer(Config.KafkaProducerClient)
	// KafkaConsumer = config.NewKafkaConsumer(Config.KafkaConsumerClient)

	KafkaProducer.SetingAsyncProducer(Config.KafkaProducerClient)

	KafkaConsumerGroup = _conf.NewKafkaConsumerGroupFromClient("myGroup", Config.KafkaConsumerClient)

	defer Config.PGORM.Close()
	defer Config.MONGO.Close()
	defer Config.KafkaProducerClient.Close()
	// defer Config.KafkaConsumerClient.Close()
	defer KafkaProducer.GetSyncProducer().Close()
	defer KafkaProducer.GetAsyncProducer().Close()
	defer KafkaConsumerGroup.GetConsumerGroup().Close()

	/* create topic each patition */
	// createTopic("users")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	middL := myMiddL.InitMiddleware()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	/* Inject Repository */

	userRepo := _user_repository.NewPsqlUserRepository(Config.PGORM)
	userMongoRepo := _user_repository.NewMongoUserRepository(Config.MONGO, _conf.DB_APPEXAMPLE)
	kafkaProRepo := _kafka_repository.NewKafkaProducerRepository(KafkaProducer)
	kafkaConRepo := _kafka_repository.NewKafkaConsumerRepository(KafkaConsumerGroup)

	/* Inject Usecase */

	kafkaUs := _kafka_usecase.NewKafkaUsecase(kafkaProRepo, kafkaConRepo)
	userUs := _user_usecase.NewUserUsecase(userRepo, userMongoRepo, kafkaUs)

	kafkaEventUs := _kafka_usecase.NewKafkaEventUsecase(userUs)

	/* Inject Handler */

	_user_handler.NewUserHandler(e, middL, userUs)
	_kafka_handler.NewKafkaHandler(e, middL, userUs)

	/* Init Consumer event handler */
	kafkaEventHandler := _kafka_event_handler.NewKafkaEventHandler(KafkaConsumerGroup.GetConsumerGroup(), kafkaEventUs)
	settingConsumerGroupHandler(kafkaEventHandler)

	/* Running web service */
	port := ":" + _conf.GetEnv("PORT", "3000")
	e.Logger.Fatal(e.Start(port))
}

func settingConsumerGroupHandler(handlerGroup sarama.ConsumerGroupHandler) {
	ctx := context.Background()
	topics := _conf.Topics
	go func(kafkaConsumerGroup *_conf.KafkaConsumerGroup) {
		if err := kafkaConsumerGroup.GetConsumerGroup().Consume(ctx, topics, handlerGroup); err != nil {
			log.Fatal(err)
		}
	}(KafkaConsumerGroup)
}
