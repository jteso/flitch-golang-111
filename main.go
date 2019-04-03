package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "1.0"

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

type HelloResponse struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := HelloResponse{
		Message: "Greetings from Golang running on Flitch!",
	}
	json.NewEncoder(w).Encode(response)
	return
}

func main() {
	httpAddr := fmt.Sprintf(":%s", getEnvOrFatal("FLITCH_AUTOBINDING_PORT"))

	log.Println("Starting server...")
	log.Printf("HTTP service listening on %s", httpAddr)

	errChan := make(chan error, 10)

	// App endpoint
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	go func() {
		errChan <- http.ListenAndServe(httpAddr, mux)
	}()

	<-errChan

	log.Println("Shutdown signal received, exiting...")
	os.Exit(1)

}

func getEnvOrFatal(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Fatalf("Environment variable not found: %s", key)
	os.Exit(1)
	return ""
}
