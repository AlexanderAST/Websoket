package handler

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
}

func newServer() *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
	}

	s.configureRouter()

	return s
}

func Start() error {

	srv := newServer()

	logrus.Info("Server started")

	return http.ListenAndServe(viper.GetString("port"), srv)
}
