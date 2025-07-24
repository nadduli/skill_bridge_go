package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nadduli/skill_bridge_go/internal/config"
	"github.com/nadduli/skill_bridge_go/internal/database"
	"github.com/nadduli/skill_bridge_go/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.HealthHandler())

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config %v", err)
	}

	db := database.ConnectDB(cfg.DatabaseURL)
	defer db.Close()

	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Server is running on http://localhost%s\n", serverAddr)

	server := &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server! %v", err)
	}

}
