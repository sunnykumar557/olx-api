package main

import (
	"log"
	"net/http"
	"time"

	"github.com/sunnykumar557/olx-api/internal/config"
	"github.com/sunnykumar557/olx-api/internal/handlers"
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

	mux.HandleFunc("GET /healthz", handlers.Healthz)

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
