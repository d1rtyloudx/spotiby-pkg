package elastic

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"net/http"
)

type Config struct {
	Addresses []string `json:"addresses"`
	Username  string   `json:"username"`
	Password  string   `env:"ELASTIC_PASSWORD"`
	APIKey    string   `env:"ELASTIC_API_KEY"`
	Header    http.Header
}

type IndexConfig struct {
	Name  string `json:"name"`
	Path  string `json:"path"`
	Alias string `json:"alias"`
}

func Connect(cfg *Config) (*elasticsearch.Client, error) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: cfg.Addresses,
		Username:  cfg.Username,
		Password:  cfg.Password,
		APIKey:    cfg.APIKey,
		Header:    cfg.Header,
	})
	if err != nil {
		return nil, fmt.Errorf("pkg.elastic.Connect - elasticsearch.NewClient: %v", err)
	}

	return client, nil
}

func MustConnect(cfg *Config) *elasticsearch.Client {
	client, err := Connect(cfg)
	if err != nil {
		panic(err)
	}

	return client
}
