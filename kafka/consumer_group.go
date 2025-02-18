package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"sync"
	"time"
)

type ConsumerGroup struct {
	brokers []string
	groupID string
	log     *zap.Logger
}

type Worker func(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int)

func (c *ConsumerGroup) ConsumeTopic(
	ctx context.Context,
	topicGroup []string,
	poolSize int,
	worker Worker,
) {
	r := NewReader(c.brokers, topicGroup, c.groupID)

	defer func() {
		if err := r.Close(); err != nil {
			c.log.Warn(
				"failed to close reader",
				zap.String("op", "pkg.kafka.ConsumeTopic - r.close()"),
				zap.Error(err),
			)
		}
	}()

	c.log.Info(
		"starting consumer group",
		zap.String("group_id", c.groupID),
		zap.Int("pool_size", poolSize),
	)

	wg := &sync.WaitGroup{}
	for i := 0; i <= poolSize; i++ {
		wg.Add(1)
		go worker(ctx, r, wg, i)
	}
	wg.Wait()
}

func NewReader(kafkaURL []string, topicGroup []string, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:                kafkaURL,
		GroupID:                groupID,
		GroupTopics:            topicGroup,
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
