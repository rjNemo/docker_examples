package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", func() http.Handler {

		type Response struct {
			Message string `json:"message"`
		}
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			response := Response{Message: "pong"}
			json.NewEncoder(w).Encode(response)
		})
	}())

	http.ListenAndServe(":80", nil) // http://0.0.0.0/
}
