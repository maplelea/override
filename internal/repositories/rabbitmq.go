package repositories

import (
	"github.com/streadway/amqp"
	"log"
	"override/config"
)

func InitRabbitMQ() (*amqp.Connection, error) {
	url := config.Viper.GetString("rabbitmq.url")
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	return conn, nil
}
