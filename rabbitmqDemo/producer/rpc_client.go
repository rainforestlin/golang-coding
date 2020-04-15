package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < 1; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func fibonacciRPC(n int) (res int, err error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", RabbitMQ.User, RabbitMQ.Password, RabbitMQ.Url, RabbitMQ.Port))
	failOnError(err, "failed to connect to rabbitmq")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"",
		true,
		false,
		true,
		false,
		nil,
	)
	failOnError(err, "failed to declare a queue")
	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to register a consumer")
	corrId := randomString(32)
	err = ch.Publish(
		"",
		"rpc_queue",
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       queue.Name,
			Body:          []byte(strconv.Itoa(n)),
		},
	)
	failOnError(err, "Failed to publish a message")
	for msg := range msgs {
		if msg.CorrelationId == corrId {
			res, err = strconv.Atoi(string(msg.Body))
			failOnError(err, "Failed to convert body to integer")
			break
		}
	}
	return
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	n := bodyFrom(os.Args)

	log.Printf(" [x] Requesting fib(%d)", n)
	res, err := fibonacciRPC(n)
	failOnError(err, "Failed to handle RPC request")

	log.Printf(" [.] Got %d", res)
}

func bodyFrom(args []string) int {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "30"
	} else {
		s = strings.Join(args[1:], " ")
	}
	n, err := strconv.Atoi(s)
	failOnError(err, "Failed to convert arg to integer")
	return n
}
