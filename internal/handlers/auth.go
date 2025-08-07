package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"sentinel/internal/auth"
	"sentinel/internal/models"
)

type AuthHandler struct {
	authService *auth.AuthService
}

func NewAuthHandler(authService *auth.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Password == "" {
		http.Error(w, `{"error":"Email and password are required"}`, http.StatusBadRequest)
		return
	}

	userResponse, err := h.authService.LoginUser(req.Email, req.Password)
	if err != nil {
		log.Printf("Login failed: %v", err)

		errorResp := map[string]string{"error": "Invalid credentials"}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(errorResp)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login successful",
		"user":    userResponse,
	})
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	userResponse, err := h.authService.RegisterUser(&req)
	if err != nil {
		log.Printf("Registration failed: %v", err)

		status := http.StatusBadRequest
		if err.Error() == "user with email "+req.Email+" already exists" {
			status = http.StatusConflict
		}

		errorResp := map[string]string{"error": err.Error()}
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errorResp)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered successfully",
		"user":    userResponse,
	})
}

func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(`{"message":"Refresh token endpoint - to be implemented"}`))
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(`{"message":"Logout endpoint - to be implemented"}`))
}

func (h *AuthHandler) OAuthCallback(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(`{"message":"OAuth callback endpoint - to be implemented"}`))
}
