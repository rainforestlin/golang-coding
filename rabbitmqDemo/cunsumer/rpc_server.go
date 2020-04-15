package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"rabbitmqDemo/config"
	"strconv"
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
		fmt.Println(n)
		return fib(n-1) + fib(n-2)
	}
}

var fibarry = []int64{0, 1, 0}

func Fib(n int64) int64 {
	for i := int64(2); i <= n; i++ {
		fibarry[2] = fibarry[0] + fibarry[1]
		fibarry[0] = fibarry[1]
		fibarry[1] = fibarry[2]
	}
	return fibarry[2]
}

func main() {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", config.RabbitMQ.User, config.RabbitMQ.Password, config.RabbitMQ.Url, config.RabbitMQ.Port))

	failOnError(err, "failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")

	defer ch.Close()
	queue, err := ch.QueueDeclare(
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

	msgs, err := ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to register a consumer")
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			n, err := strconv.Atoi(string(msg.Body))
			failOnError(err, "failed to convert to integer")
			log.Printf("[ .]fib(%d)", n)
			response := Fib(int64(n))
			err = ch.Publish(
				"",
				msg.ReplyTo,
				false,
				false,
				amqp.Publishing{ContentType: "text/plain",
					CorrelationId: msg.CorrelationId,
					Body: []byte(strconv.FormatInt(response, 10),
					)},
			)
			failOnError(err, "failed to publish a message")
			msg.Ack(false)
		}
	}()
	log.Printf(" [*]Awaiting RPC requests")
	<-forever
}
