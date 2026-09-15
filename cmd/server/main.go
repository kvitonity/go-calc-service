package main

import (
	"log"
	"net/http"
	"time"

	"github.com/kvitonity/go-calc-service/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/expression", handler.ExpressionHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Println("Server starting on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
