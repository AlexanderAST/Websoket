package handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type Config struct {
	BindAddr string `yaml:"bind_addr"`
	LogLevel string `yaml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}

func Start(config *Config) error {

	srv := newServer()

	logrus.Info("Server started")

	return http.ListenAndServe(config.BindAddr, srv)
}
