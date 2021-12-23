package main

import (
	"fmt"
	"log"

	"github.com/hyperstone1/Rabbitmq_go/repository"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("% s:% s", msg, err)
	}
}

func main() {
	fmt.Println("Consumer Application")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()
	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	rep := repository.New()
	forever := make(chan struct{})
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
			rep.Record(string(d.Body))
		}
	}()

	fmt.Println("Successfully connected to RabbitMQ instance")
	fmt.Println(" [*] - waiting for messages")

	<-forever
}
