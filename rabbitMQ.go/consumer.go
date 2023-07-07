package rabbitMQ

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

func RabbitMQConsumer(){
	
	fmt.Println("ConsumeSocket");

	address := os.Getenv("RABBITMQ_URL")

	conn,err := amqp.Dial(address)
	failOnError(err,"Connection failed")
	defer conn.Close()

	ch,err := conn.Channel()
	failOnError(err,"Channel failed")
	defer ch.Close()

	msgs, err := ch.Consume("TestQueue", "", true, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")
	forever := make(chan bool)

	go func(){
		for d := range msgs {
			fmt.Printf("Received a message: %s\n", d.Body)
		}
	}()

	fmt.Println("===============")
	fmt.Println(" [*] - waiting for messages")
	<-forever	

}