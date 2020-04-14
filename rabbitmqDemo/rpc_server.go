package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"rabbitmqDemo/config"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1

	} else {
		return fib(n-1) * fib(n-2)
	}
}

func main() {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", config.RabbitMQ.User, config.RabbitMQ.Password, config.RabbitMQ.Url, config.RabbitMQ.Port))

	failOnError(err, "failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")

	defer ch.Close()
	_,err = ch.QueueDeclare(
		"rpc_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	err = ch.Qos(
		1,
		0,
		false,
	)
	failOnError(err, "Failed to set Qos")
}
