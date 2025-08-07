package main

import (
	"log"
	"net/http"

	"sentinel/internal/config"
	"sentinel/internal/handlers"
	"sentinel/internal/middleware"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	cfg := config.Load()

	r := mux.NewRouter()
	
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.CORSMiddleware)

	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(middleware.JSONMiddleware)

	api.HandleFunc("/health", handlers.HealthHandler).Methods("GET")

	authHandler := handlers.NewAuthHandler()
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", authHandler.Login).Methods("POST")
	auth.HandleFunc("/register", authHandler.Register).Methods("POST")
	auth.HandleFunc("/refresh", authHandler.RefreshToken).Methods("POST")
	auth.HandleFunc("/logout", authHandler.Logout).Methods("POST")
	auth.HandleFunc("/oauth/callback", authHandler.OAuthCallback).Methods("GET")

	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}