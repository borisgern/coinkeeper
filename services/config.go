package services

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Logger *LoggerConfig `yaml:"logger"`
	HTTPServer *HTTPServerConfig `yaml:"http_server"`
	GRPCServer *GRPCServerConfig `yaml:"grpc_server"`
}

type LoggerConfig struct {
	LoggerLevel string `yaml:"log_level"`
}

type HTTPServerConfig struct {
	ServerPort     string `yaml:"server_port"`
	ServerHost     string `yaml:"server_host"`
	GinReleaseMode bool   `yaml:"gin_release_mode"`
}

type GRPCServerConfig struct {
	ServerPort     string `yaml:"server_port"`
	ServerHost     string `yaml:"server_host"`
	GinReleaseMode bool   `yaml:"gin_release_mode"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	SslMode  string `yaml:"ssl_mode"`
	User     string `yaml:"user"`
	Port     string `yaml:"port"`
}

func NewConfig() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return &Config{}, fmt.Errorf("CONFIG_PATH env must be defined")
	}

	if configBytes, err := ioutil.ReadFile(configPath); err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	} else {
		var c Config
		if err = yaml.Unmarshal(configBytes, &c); err != nil {
			return nil, fmt.Errorf("failed to unmarshal yaml config: %v", err)
		}
		return &c, nil
	}
}