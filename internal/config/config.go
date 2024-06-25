package config

import "time"

type Config struct {
	Env         string     `yaml:"env"          env-default:"local"`
	StoragePath string     `yaml:"storage_path"                     env-required:"true"`
	GRPC        GRPCConfig `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}
