package consumer

import (
	"lebrancconvas/gokafka/config"
	"lebrancconvas/gokafka/pkg/utils"
)

func main() {
	// Kafka Connection
	cfg := config.KafkaConfig{
		Url: "localhost:9092",
		Topic: "player",
	}
	connection := utils.KafkaConnect(cfg)
	defer connection.Close()

	// Message Reading
}
