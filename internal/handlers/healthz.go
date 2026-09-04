package handlers

import "net/http"

func Healthz(w http.ResponseWriter, r *http.Request) {
	// Order is very important here.
	// First, set the content type,
	// then write the status code,
	// and finally write the response body.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(`{"status":"okay!"}`))
}
