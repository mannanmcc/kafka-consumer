package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func StartConsumer() {
	config := kafka.ReaderConfig{
		Brokers:  []string{"localhost: 9092"},
		Topic:    "first_topic",
		MaxBytes: 10,
	}

	reader := kafka.NewReader(config)

	//loop to make sure it does exit after reading a single message
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("error reading kafka message", err)
			//do not exist after encounter a single message
			continue
		}

		fmt.Println("message content::", string(m.Value))
	}
}
