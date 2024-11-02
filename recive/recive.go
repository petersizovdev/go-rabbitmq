package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("fail to connect RMQ", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("fail to open ch", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to declare a queue", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to register a consumer", err)
	}

	var forever chan struct {
	}
	go func() {
		for d := range msgs {
			log.Printf("Recived a msg: %s", d.Body)
		}
	}()
	log.Printf("Waiting for msgs...")
	<-forever

}
