package main

import (
	"log"
	"net/http"
	"time"

	"github.com/sunnykumar557/olx-api/internal/config"
)

func main() {
	cfg := config.MustLoad()

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

		w.Write([]byte(`{"status":"okay!"}`))
	})

	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}
	log.Printf("server is running on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
