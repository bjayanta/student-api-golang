package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bjayanta/students-api/internal/config"
)

func main() {
	// fmt.Println("Welcome to students api")

	// Load config
	cfg := config.MustLoad()

	// Database setup

	// Setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func (w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte("Welcome to students api"))
	})

	// Setup server
	server := http.Server{
		Addr: cfg.Address,
		Handler: router,
	}

	fmt.Printf("Server started at %s", cfg.HTTPServer.Address)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server")
	}

}