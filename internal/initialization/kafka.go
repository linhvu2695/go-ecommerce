package initialization

import (
	"go-ecommerce/global"
	"log"

	"github.com/segmentio/kafka-go"
)

var KafkaProducer *kafka.Writer

func InitKafka() {
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP("localhost:19092"),
		Topic:    "register_topic",
		Balancer: &kafka.LeastBytes{},
	}
}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		log.Fatalf("Error closing Kafka producer: %v", err)
	}
}
