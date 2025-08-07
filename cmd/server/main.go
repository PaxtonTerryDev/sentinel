package main

import (
	"log"
	"net/http"

	"sentinel/internal/auth"
	"sentinel/internal/config"
	"sentinel/internal/handlers"
	"sentinel/internal/middleware"
	"sentinel/pkg/database"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	cfg := config.Load()

	db, err := database.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	userRepo := auth.NewUserRepository(db)
	authService := auth.NewAuthService(userRepo)

	r := mux.NewRouter()

	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.CORSMiddleware)

	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(middleware.JSONMiddleware)

	api.HandleFunc("/health", handlers.HealthHandler).Methods("GET")

	authHandler := handlers.NewAuthHandler(authService)
	authRoutes := api.PathPrefix("/auth").Subrouter()
	authRoutes.HandleFunc("/login", authHandler.Login).Methods("POST")
	authRoutes.HandleFunc("/register", authHandler.Register).Methods("POST")
	authRoutes.HandleFunc("/refresh", authHandler.RefreshToken).Methods("POST")
	authRoutes.HandleFunc("/logout", authHandler.Logout).Methods("POST")
	authRoutes.HandleFunc("/oauth/callback", authHandler.OAuthCallback).Methods("GET")

	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
