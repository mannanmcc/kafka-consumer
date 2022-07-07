package kafka

import (
	"context"
	"fmt"

	conf "github.com/mannanmcc/kafka-consumer/internal/config"
	"github.com/segmentio/kafka-go"
)

func StartConsumer(conf *conf.Kafka) {
	config := kafka.ReaderConfig{
		Brokers:  []string{conf.Host + ":" + conf.Port},
		Topic:    conf.Topic,
		MaxBytes: conf.MaxBytes,
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
