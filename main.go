package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Print("Savepoint API is starting...")
	runServer()
}

func runServer() {
	type JsonResponse struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}

	response := JsonResponse{
		Message: "Welcome to SavePoint API",
		Status:  "running",
	}

	//listen for request on path /bar
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// start server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
