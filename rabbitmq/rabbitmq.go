package rabbitmq

import (
	"context"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"golang.org/x/sync/errgroup"
)

type ExchangeAndQueueBinding struct {
	ExchangeName string `yaml:"exchange_name"`
	ExchangeKind string `yaml:"exchange_kind"`
	RoutingKey   string `yaml:"routing_key"`
	QueueName    string `yaml:"queue_name"`
	Concurrency  int    `yaml:"concurrency"`
	ConsumerTag  string `yaml:"consumer_tag"`
}

type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `env:"RABBITMQ_PASSWORD"`
}

func NewRabbitMQConn(cfg *Config) (*amqp091.Connection, error) {
	amqpConn, err := amqp091.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d", cfg.Username, cfg.Password, cfg.Host, cfg.Port))
	if err != nil {
		return nil, fmt.Errorf("pkg.rabbitmq.NewRabbitMQConn = amqp091.Dial: %w", err)
	}

	return amqpConn, nil
}

func DeclareQueue(amqpChan *amqp091.Channel, name string) (amqp091.Queue, error) {
	queue, err := amqpChan.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return amqp091.Queue{}, fmt.Errorf("pkg.rabbitmq.DeclareQueue - amqpChan.QueueDeclare: %w", err)
	}

	return queue, nil
}

func DeclareExchange(amqpChan *amqp091.Channel, name string, kind string) error {
	err := amqpChan.ExchangeDeclare(
		name,
		kind,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("pkg.rabbitmq.DeclareExchange - amqpChan.ExchangeDeclare: %w", err)
	}

	return nil
}

func BindExchangeAndQueue(amqpChan *amqp091.Channel, exchangeName string, queueName string, key string) error {
	err := amqpChan.QueueBind(queueName, key, exchangeName, false, nil)
	if err != nil {
		return fmt.Errorf("pkg.rabbitmq.BindExchangeAndQueue - amqpChan.QueueBind: %w", err)
	}

	return nil
}

type ConsumeDeliveries func(ctx context.Context, deliveries <-chan amqp091.Delivery, workerID int) func() error

func ConsumeQueue(
	ctx context.Context,
	amqpChan *amqp091.Channel,
	concurrency int,
	queue string,
	consumer string,
	worker ConsumeDeliveries,
) error {
	deliveries, err := amqpChan.Consume(
		queue,
		consumer,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("pkg.rabbitmq.ConsumeQueue - amqpChan.Consume: %w", err)
	}

	eg, ctx := errgroup.WithContext(ctx)
	for i := 0; i < concurrency; i++ {
		eg.Go(worker(ctx, deliveries, i))
	}

	return eg.Wait()
}
