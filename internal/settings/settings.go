package settings

import (
	"context"
)

type Config struct {
	Recievers []Reciever `yaml:"receievers"`
}

type Reciever struct {
	Address     string `yaml:"address"`
	ServiceType string `yaml:"service_type"`
}

type ConfigReader interface {
	Load(context.Context) (Config, error)
}
