package config

import (
	"github.com/caarlos0/env"
)

type Config struct {
	AppPort string `env:"APP_PORT" envDefault:":8080"`

	KafkaServer  []string `env:"KAFKA_SERVER" envDefault:"localhost:9092" envSeparator:","`
	KafkaGroup   string   `env:"KAFKA_GROUP" envDefault:"Messages"`
	MessageTopic string   `env:"MESSAGE_TOPIC" envDefault:"messagetopic"`
}

func Get() *Config {
	config := &Config{}
	err := env.Parse(config)
	if err != nil {
		panic(err)
	}
	return config
}
