package main

import (
	"fmt"
	
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	APIHost string `envconfig:"API_HOST" default:"0.0.0.0:9333"`
	PgHost string `envconfig:"PG_HOST" default:"postgres://outbox:outbox@localhost:5432/outbox?sslmode=disable"`
}

func LoadConfig() (*Config, error) {
	var config Config

	if err := envconfig.Process("app", &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	config, err := LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	
	api, err := NewAPI(config.APIHost, config.PgHost)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	if err := api.Start(); err != nil {
		fmt.Println(err)
		return
	}
}

