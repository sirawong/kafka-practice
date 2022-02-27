package main

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"consumer/config"
	repository "consumer/repository/database"
	"consumer/service"
)

func main() {
	appConfig := config.Get()
	ctx := context.Background()

	consumer := initKafka(appConfig)
	defer consumer.Close()

	db := initDatabase(ctx, appConfig)
	defer db.Disconnect(ctx)

	dbRepo := repository.New(db, appConfig)
	srvHandlerData := service.NewHandleData(dbRepo, appConfig)
	srvConsumer := service.NewConsumerHandler(srvHandlerData)

	fmt.Println("consumer start...")
	for {
		consumer.Consume(ctx, []string{appConfig.MessageTopic}, srvConsumer)
	}
}

func initKafka(config *config.Config) (sarama.ConsumerGroup)  {
	saraConfig := sarama.NewConfig()
	saraConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	consumer, err := sarama.NewConsumerGroup(config.KafkaServer, config.KafkaGroup, saraConfig)
	if err != nil {
		panic(err)
	}
	return consumer
}

func initDatabase(ctx context.Context, config *config.Config) *mongo.Client {
	dsn := fmt.Sprintf("%v", config.MongoEndpoint)
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		panic(err)
	}
	return db
}