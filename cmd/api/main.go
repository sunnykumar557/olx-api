package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	mux := http.NewServeMux()

	// http.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
	// 	//w.Header().Add("content-type", "application/json")
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte(`{"status":"ok"}`))
	// })

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		// Order is very important here.
		// First, set the content type,
		// then write the status code,
		// and finally write the response body.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(`{"status":"ok"}`))
	})

	srv := http.Server{
		Addr:         ":8090",
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
