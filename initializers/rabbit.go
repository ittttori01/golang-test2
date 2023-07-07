package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}


func ConnectRabbitMq(){

	address := os.Getenv("RABBITMQ_URL")
	
	conn,err := amqp.Dial(address)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
        "TestQueue", // name
        false, // durable
        false, // delete when unused
        false,
		false, // exclusive
		nil, // exclusive)
	)
    failOnError(err, "Failed to declare a queue")
	fmt.Println(q)

	err = ch.Publish("","TestQueue", false, false, amqp.Publishing{
		ContentType: "text/plain",
        Body:        []byte("hello QUEUEUEUEUEU RABBIT"),
	})
	failOnError(err, "Failed to publish a message")

	fmt.Println("RabbitMQ connection established")
}