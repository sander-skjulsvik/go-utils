package rest

import "net/http"

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte{})
}

func MyMux() *http.ServeMux {
	// Init request multiplexer
	mux := http.NewServeMux()
	// Adding health check handler
	mux.HandleFunc("/health", HealthCheckHandler)
	return mux
}
