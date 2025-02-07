package rabbitmq

import (
	"context"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
)

type PublisherConfig struct {
	ExchangeName string `yaml:"exchange"`
	RoutingKey   string `yaml:"routing_key"`
}

type Publisher struct {
	amqpConn *amqp091.Connection
	amqpChan *amqp091.Channel
}

func (p *Publisher) Close() error {
	if err := p.amqpConn.Close(); err != nil {
		return fmt.Errorf("pkg.rabbitmq.publisher.Close - p.amqpConn.Close: %w", err)
	}

	if err := p.amqpChan.Close(); err != nil {
		return fmt.Errorf("pkg.rabbitmq.publisher.Close - p.amqpChan.Close: %w", err)
	}

	return nil
}

func CreatePublisher(cfg *Config) (*Publisher, error) {
	amqpConn, err := NewRabbitMQConn(cfg)
	if err != nil {
		return nil, err
	}

	amqpChan, err := amqpConn.Channel()
	if err != nil {
		return nil, fmt.Errorf("pkg.rabbitmq.publisher.CreatePublsiher - amqpConn.Channel: %w", err)
	}

	return &Publisher{
		amqpConn: amqpConn,
		amqpChan: amqpChan,
	}, nil
}

func MustCreatePublisher(cfg *Config) *Publisher {
	publisher, err := CreatePublisher(cfg)
	if err != nil {
		panic(err)
	}

	return publisher
}

func (p *Publisher) PublishWithContext(
	ctx context.Context,
	exchange string,
	key string,
	mandatory bool,
	immediate bool,
	msg amqp091.Publishing,
) error {
	err := p.amqpChan.PublishWithContext(ctx, exchange, key, mandatory, immediate, msg)
	if err != nil {
		return fmt.Errorf("pkg.rabbitmq.PublishWithContext - p.amqpChan.PublishWithContext: %w", err)
	}

	return nil
}
