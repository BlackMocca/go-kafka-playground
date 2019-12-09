package main

import (
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
)

var (
	Config              *_conf.Config
	KafkaProducerClient sarama.Client
	KafkaConsumerClient sarama.Client
	KafkaProducer       *_conf.KafkaProducer
	KafkaConsumer       *_conf.KafkaConsumer
)

func init() {
	Config = _conf.NewConfigWithService(_conf.NewPsqlConnection(), _conf.NewMongoSession)
	KafkaProducerClient, KafkaConsumerClient = _conf.NewKafkaClient()
	Config.KafkaProducerClient = KafkaConsumerClient
	Config.KafkaConsumerClient = KafkaConsumerClient
}

func main() {
	KafkaProducer = config.NewKafkaProducer(Config.KafkaProducerClient)
	KafkaConsumer = config.NewKafkaConsumer(Config.KafkaConsumerClient)

	defer Config.PGORM.Close()
	defer Config.MONGO.Close()
	defer Config.KafkaProducerClient.Close()
	defer Config.KafkaConsumerClient.Close()

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

	/* Inject Usecase */

	userUs := _user_usecase.NewUserUsecase(userRepo)

	/* Inject Handler */

	_user_handler.NewUserHandler(e, middL, userUs)

	port := ":" + _conf.GetEnv("PORT", "3000")
	e.Logger.Fatal(e.Start(port))
}
