package main

import (
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"

	"producer/config"
	"producer/handler"
	repository "producer/repository/kafka"
	service "producer/service"
)

func main() {
	appConfig := config.Get()
	producer := initKafkaProducer(appConfig)
	defer producer.Close()

	kafkaRepo := repository.New(producer, appConfig)
	senderSrv := service.New(kafkaRepo)

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	_ = handler.New(senderSrv).RegisterRoutes(router)
	router.Run(appConfig.AppPort)
}

func initKafkaProducer(appConfig *config.Config) sarama.SyncProducer {
	producer, err := sarama.NewSyncProducer(appConfig.KafkaServer, nil)
	if err != nil {
		panic(err)
	}
	return producer
}
