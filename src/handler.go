package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func buildServerMux(server *http.ServeMux) *http.ServeMux {
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		responseBody := map[string]string{
			"message": "Hello World!",
		}
		json.NewEncoder(w).Encode(responseBody)
	})

	server.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request to /slow")
		w.Header().Set("Content-Type", "application/json")

		responseBody := map[string]string{
			"message": "This is a slow response",
		}
		json.NewEncoder(w).Encode(responseBody)

		// Simulate a slow response
		time.Sleep(10 * time.Second)
	})

	return server
}
