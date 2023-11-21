package handler

import (
	"encoding/json"
	"net/http"
)

func (s *server) helloFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode("hello")
		if err != nil {
			return
		}
	}
}
