package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `env:"REDIS_PASSWORD"`
	DB       int    `yaml:"db"`
}

func Connect(cfg *Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, fmt.Errorf("redis.Connect - client.Ping: %w", err)
	}

	return client, nil
}

func MustConnect(cfg *Config) *redis.Client {
	client, err := Connect(cfg)
	if err != nil {
		panic(err)
	}

	return client
}
