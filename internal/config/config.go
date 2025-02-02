package config

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

type DatabaseConfig struct {
        ConnectionString string `yaml:"connectionString"`
    }

type Config struct {
    AllowedOrigins []string `yaml:"allowedOrigins"`
    BaseUrl string `yaml:"baseUrl"`
    Database DatabaseConfig `yaml:"database"`
}

func LoadConfig(reader io.Reader) (*Config, error) {
    config := &Config{
        AllowedOrigins: []string{"localhost"},
        BaseUrl: "localhost:8081",
        Database: DatabaseConfig{ ConnectionString: "localhost", },
    }

    b, err := io.ReadAll(reader)
    if err != nil {
        return nil, fmt.Errorf("Failed to read config: %w", err)
    }

    err = yaml.Unmarshal(b, config)
    if err != nil {
        return nil, err
    }

    return config, nil
}
