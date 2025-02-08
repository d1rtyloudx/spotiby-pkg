package minio

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Config struct {
	Port            int    `yaml:"port"`
	Host            string `yaml:"host"`
	UseSSL          bool   `yaml:"use_ssl"`
	AccessKeyID     string `yaml:"access_key_id"`
	SecretAccessKey string `env:"MINIO_SECRET"`
}

func Connect(cfg *Config) (*minio.Client, error) {
	client, err := minio.New(
		fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		&minio.Options{
			Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
			Secure: cfg.UseSSL,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("pkg.minio.Connect - minio.New: %w", err)
	}

	return client, nil
}

func MustConnect(cfg *Config) *minio.Client {
	client, err := Connect(cfg)
	if err != nil {
		panic(err)
	}

	return client
}
