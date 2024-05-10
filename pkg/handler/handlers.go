package handler

import (
	"net/http"
)

func (s *server) helloFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		s.respond(w, r, http.StatusOK, map[string]interface{}{"hello": "world"})
	}
}
