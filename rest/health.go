package rest

import "net/http"

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte{})
}
