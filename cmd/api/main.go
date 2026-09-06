package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sunnykumar557/olx-api/internal/config"
	"github.com/sunnykumar557/olx-api/internal/db"
	"github.com/sunnykumar557/olx-api/internal/handlers"
)

func main() {
	cfg := config.MustLoad()
	db, err := db.Connect(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("main.db.connect: %v", err)
	}

	fmt.Println("DB connected successfully")
	fmt.Println("starting olx server...")

	mux := http.NewServeMux()

	// Endpoints
	mux.HandleFunc("GET /healthz", handlers.Healthz)
	mux.HandleFunc("GET /listings", handlers.List(db))
	mux.HandleFunc("DELETE /listings/{id}", handlers.DeleteListing(db))

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
