package main

import (
	"fmt"
	"time"

	conf "github.com/mannanmcc/kafka-consumer/internal/config"
	appkafka "github.com/mannanmcc/kafka-consumer/internal/kafka"
)

func main() {
	fmt.Println("starting kafka consumer.....")
	//config := conf.NewConfig()
	//run the consumer in the separate go routing to make sure it does not stop on this line as we need to execute rest of the code
	go appkafka.StartConsumer(conf.GetConf().Kafka)

	fmt.Println("consumer started")

	//keep it running for 10  mins
	time.Sleep(10 * time.Minute)
}
