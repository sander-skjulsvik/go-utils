package main

import (
	"log"
	"net/http"

	"github.com/sander-skjulsvik/go-utils/rest"
)

func main() {
	addr := ":3000"
	// Init request multiplexer
	mux := http.NewServeMux()

	// Handlers
	mux.HandleFunc("/health", rest.HealthCheckHandler)

	// Serve
	log.Printf("Listening: %s", addr)
	http.ListenAndServe(addr, mux)

}
