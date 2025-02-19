package elastic

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"net/http"
)

type Config struct {
	Addresses []string `json:"addresses"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	APIKey    string   `json:"api_key"`
	Header    http.Header
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
