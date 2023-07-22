package utils

import (
	"context"
	"lebrancconvas/gokafka/config"

	"github.com/segmentio/kafka-go"
)

func KafkaConnect(cfg config.KafkaConfig) (*kafka.Conn) {
	connection, err := kafka.DialLeader(context.Background(), "tcp", cfg.Url, cfg.Topic, 0)
	if err != nil {
		panic(err.Error())
	}
	return connection
}

func CheckTopicExists(connection *kafka.Conn, topic string) bool {
	partitions, err := connection.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	for _, p := range partitions {
		if p.Topic == topic {
			return true
		}
	}
	return false
}
