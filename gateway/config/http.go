package config

import "os"

type HTTPConfig struct {
	Host string
	Port string
}

func LoadHTTPConfig() HTTPConfig {
	return HTTPConfig{
		Host: os.Getenv("HTTP_HOST"),
		Port: os.Getenv("HTTP_PORT"),
	}
}
