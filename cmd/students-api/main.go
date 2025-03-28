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
	"github.com/bjayanta/students-api/internal/http/handlers/student"
	"github.com/bjayanta/students-api/internal/storage/sqlite"
)

func main() {
	// fmt.Println("Welcome to students api")

	// Load config
	cfg := config.MustLoad()

	// Database setup
	_, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Connected to database", slog.String("env", cfg.Env), slog.String("version", "1.0.0"), slog.String("storage", cfg.StoragePath))

	// Setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

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