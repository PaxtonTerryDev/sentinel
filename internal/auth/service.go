package auth

import (
	"fmt"
	"strings"
	"time"

	"sentinel/internal/auth/password"
	"sentinel/internal/models"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// AuthService handles authentication business logic and depends on the UserRepository interface.
// This design allows the service to work with any implementation of UserRepository
// (PostgreSQL, Redis, MongoDB, Mock, etc.) without knowing the storage details.
//
// The service layer focuses on:
// - Business logic and validation
// - Password hashing/verification
// - Error handling and logging
// - Data transformation between request/response models
//
// Example usage with different repository implementations:
//   postgresRepo := NewUserRepository(postgresDB)
//   redisRepo := NewRedisUserRepository(redisClient)
//   mockRepo := &MockUserRepository{}
//   
//   // All work with the same service
//   authService := NewAuthService(postgresRepo) // or redisRepo, mockRepo
type AuthService struct {
	userRepo  UserRepository
	validator *validator.Validate
}

// NewAuthService creates a new AuthService with the provided UserRepository implementation.
// The repository implementation is injected via dependency injection, following the
// Dependency Inversion Principle (high-level modules should not depend on low-level modules).
func NewAuthService(userRepo UserRepository) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		validator: validator.New(),
	}
}

func (s *AuthService) RegisterUser(req *models.CreateUserRequest) (*models.UserResponse, error) {
	if err := s.validateRegistrationRequest(req); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	normalizedEmail := strings.ToLower(strings.TrimSpace(req.Email))

	exists, err := s.userRepo.Exists(normalizedEmail)
	if err != nil {
		return nil, fmt.Errorf("failed to check if user exists: %w", err)
	}

	if exists {
		return nil, fmt.Errorf("user with email %s already exists", normalizedEmail)
	}

	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	now := time.Now()
	user := &models.User{
		ID:           uuid.New(),
		Email:        normalizedEmail,
		PasswordHash: hashedPassword,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		IsActive:     true,
		IsVerified:   false,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &models.UserResponse{
		ID:         user.ID,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		IsActive:   user.IsActive,
		IsVerified: user.IsVerified,
		CreatedAt:  user.CreatedAt,
		LastLogin:  user.LastLogin,
	}, nil
}

func (s *AuthService) LoginUser(email, plainPassword string) (*models.UserResponse, error) {
	normalizedEmail := strings.ToLower(strings.TrimSpace(email))
	
	user, err := s.userRepo.GetByEmail(normalizedEmail)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if !user.IsActive {
		return nil, fmt.Errorf("user account is deactivated")
	}

	if err := password.VerifyPassword(user.PasswordHash, plainPassword); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if err := s.userRepo.UpdateLastLogin(user.ID); err != nil {
		return nil, fmt.Errorf("failed to update last login: %w", err)
	}

	return &models.UserResponse{
		ID:         user.ID,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		IsActive:   user.IsActive,
		IsVerified: user.IsVerified,
		CreatedAt:  user.CreatedAt,
		LastLogin:  &time.Time{},
	}, nil
}

func (s *AuthService) validateRegistrationRequest(req *models.CreateUserRequest) error {
	if err := s.validator.Struct(req); err != nil {
		return err
	}

	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	if len(req.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	return nil
}