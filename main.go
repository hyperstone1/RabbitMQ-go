package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("% s:% s", msg, err)
	}
}

func main() {
	fmt.Println("Welcome to RabbitMQ")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	fmt.Println("Successfully Connection")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	fmt.Println(q)

	for i := 0; i < 2000; i++ {

		err = ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte("Hello World"),
			})
	}
	failOnError(err, "Failed to publish a message")

	fmt.Println("Successfully Published Message to Queue")
}
