package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `env:"POSTGRES_PASSWORD"`
	DBName   string `yaml:"db_name"`
	SSLMode  string `yaml:"ssl_mode"`
}

func Connect(cfg *Config) (*sqlx.DB, error) {
	dataSrcName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := sqlx.Connect("postgres", dataSrcName)
	if err != nil {
		return nil, fmt.Errorf("pkg.postgres.Connect - sqlx.Connect: %w", err)
	}

	return db, nil
}

func MustConnect(cfg *Config) *sqlx.DB {
	db, err := Connect(cfg)
	if err != nil {
		panic(err)
	}

	return db
}
