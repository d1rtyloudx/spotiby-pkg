package kafka

import (
	"fmt"
	"github.com/segmentio/kafka-go"
)

type Config struct {
	Brokers []string `yaml:"brokers"`
	GroupID string   `yaml:"group_id"`
}

type TopicConfig struct {
	TopicName         string `yaml:"topic_name"`
	Partitions        int    `yaml:"partitions"`
	ReplicationFactor int    `yaml:"replication_factor"`
}

func NewKafkaConn(addr string) (*kafka.Conn, error) {
	kafkaConn, err := kafka.Dial("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("pkg.kafka.NewKafkaConn - kafka.Dial: %w", err)
	}

	return kafkaConn, nil
}
