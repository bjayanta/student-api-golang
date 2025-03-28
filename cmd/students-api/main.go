package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	slog.Info("Server started at", slog.String("address", cfg.Address))
	// fmt.Printf("Server started at %s", cfg.HTTPServer.Address)

	// Create chanel
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	

	// Start server
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	} ()

	<-done

	slog.Info("Shutting down the server...")
	
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	// Shutdown server
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully!")

}