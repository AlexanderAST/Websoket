package handler

import (
	"net/http"
)

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {

	s.router.HandleFunc("/hello", s.helloFunc()).Methods("GET")
}
