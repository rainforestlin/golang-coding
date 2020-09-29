package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = "topic_demo"

	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	defer client.Close()
	if err != nil {
		fmt.Println("producer closed, err :", err)
		return
	}
	x := 0
	for {
		x += 1
		msg.Value = sarama.StringEncoder(fmt.Sprintf("test %d start", x))
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send message err: ", err)
			return
		}
		fmt.Printf("pid: %v, offset:%v \n", pid, offset)
		time.Sleep(1 * time.Second)
	}

}
