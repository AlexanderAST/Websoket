package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	Upserver "gorillaWebsoket/pkg/handler"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Cannot init config: %s", err.Error())
	}

	if err := Upserver.Start(); err != nil {
		logrus.Fatalf("cannot start server %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
