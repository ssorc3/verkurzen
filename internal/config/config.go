package config

import (
	"encoding/json"
	"log"

	"github.com/spf13/viper"
)

var viperInstance = viper.New()
var Default Config

type Config struct {
    Server struct {
        Port uint
        Host string
    }
    Database struct {
        URL string
    }
}

func (d Config) String() string {
    b, _ := json.Marshal(d)
    return string(b)
}

func Parse() Config {
    if err := viperInstance.Unmarshal(&Default); err != nil {
        log.Fatal("Failed to read configuration", err)
    }
    return Default
}

func Viper() *viper.Viper {
    return viperInstance
}
