package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	w *kafka.Writer
}

func (p *Producer) PublishMessage(ctx context.Context, msgs ...kafka.Message) error {
	return p.w.WriteMessages(ctx, msgs...)
}

func (p *Producer) Close() error {
	return p.w.Close()
}
