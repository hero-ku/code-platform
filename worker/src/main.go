package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	channel, err := conn.Channel()
	failOnError(err, "Failed to get channel")
	defer channel.Close()

	queue, err := channel.QueueDeclare("Programs", true, false, false, false,
		amqp.Table{
			amqp.QueueTypeArg: amqp.QueueTypeClassic,
		},
	)
	failOnError(err, "Failed to create queue")

	msgs, err := channel.Consume(queue.Name, "", true, false, false, false, nil)

	for delivery := range msgs {
		log.Printf("%s", delivery.Body)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s:%s", msg, err)
	}
}
