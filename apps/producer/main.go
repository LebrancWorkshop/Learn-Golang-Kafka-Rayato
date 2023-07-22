package main

import (
	"lebrancconvas/gokafka/config"
	"lebrancconvas/gokafka/models"
	"lebrancconvas/gokafka/pkg/utils"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Kafka Connection
	cfg := config.KafkaConfig{
		Url: "localhost:9092",
		Topic: "player",
	}
	connection := utils.KafkaConnect(cfg)
	defer connection.Close()

	// Check Topic is exists. If not exists, create topic
	if !utils.CheckTopicExists(connection, cfg.Topic) {
		topicConfigs := []kafka.TopicConfig{
			{
				Topic: cfg.Topic,
				NumPartitions: 1,
				ReplicationFactor: 1,
			},
		}

		err := connection.CreateTopics(topicConfigs...)
		if err != nil {
			panic(err.Error())
		}
	}

	// Assign Message to produce.

	data := func () ([]kafka.Message) {
		players := []models.Player{
			{
				Id: 1,
				Name: "Bangbua",
			},
			{
				Id: 2,
				Name: "Kraton",
			},
		}

		// Convert Message to JSON Byte.
		messages := make([]kafka.Message, 0)
		for _, player := range players {
			messages = append(messages, kafka.Message{
				Value: utils.CompressToJSONByte(&player),
			})
		}

		return messages
	}


	// Set Timeout

	// Close Connection
}
