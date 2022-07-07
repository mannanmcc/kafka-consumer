package main

import (
	"fmt"
	"time"

	appkafka "github.com/mannanmcc/kafka-consumer/internal/kafka"
)

func main() {
	fmt.Println("starting kafka consumer.....")

	//run the consumer in the separate go routing to make sure it does not stop on this line as we need to execute rest of the code
	go appkafka.StartConsumer()

	fmt.Println("consumer started")

	//keep it running for 10  mins
	time.Sleep(10 * time.Minute)
}
