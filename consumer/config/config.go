package config

import (
	"github.com/caarlos0/env"
)

type Config struct {
	KafkaServer  []string `env:"KAFKA_SERVER" envDefault:"localhost:9092" envSeparator:","`
	KafkaGroup   string   `env:"KAFKA_GROUP" envDefault:"Messages"`
	MessageTopic string   `env:"MESSAGE_TOPIC" envDefault:"messagetopic"`

	MongoEndpoint     string `env:"MONGODB_ENDPOINT" envDefault:"mongodb://localhost:27017"`
	MongoDBName       string `env:"MONGODB_DBNAME" envDefault:"messagesDB"`
	MongoDBCollection string `env:"MONGODB_COLLECTION" envDefault:"messagesColl"`
}

func Get() *Config {
	config := &Config{}
	err := env.Parse(config)
	if err != nil {
		panic(err)
	}
	return config
}
