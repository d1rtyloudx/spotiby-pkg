package kafka

import (
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
	"time"
)

type Config struct {
	Brokers []string `yaml:"brokers"`
	GroupID string   `yaml:"groupID"`
}

type TopicConfig struct {
	TopicName         string `yaml:"topicName"`
	Partitions        int    `yaml:"partitions"`
	ReplicationFactor int    `yaml:"replicationFactor"`
}

func NewReader(kafkaURL []string, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:                kafkaURL,
		GroupID:                groupID,
		Topic:                  topic,
		MinBytes:               minBytes,
		MaxBytes:               maxBytes,
		QueueCapacity:          queueCapacity,
		HeartbeatInterval:      heartbeatInterval,
		CommitInterval:         commitInterval,
		PartitionWatchInterval: partitionWatchInterval,
		MaxAttempts:            maxAttempts,
		MaxWait:                time.Second,
		Dialer: &kafka.Dialer{
			Timeout: dialTimeout,
		},
	})
}

func NewWriter(brokers []string) *kafka.Writer {
	w := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: writerRequiredAcks,
		MaxAttempts:  writerMaxAttempts,
		Compression:  compress.Snappy,
		ReadTimeout:  writerReadTimeout,
		WriteTimeout: writerWriteTimeout,
		Async:        false,
	}
	return w
}
