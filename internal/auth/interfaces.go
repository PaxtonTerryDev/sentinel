package auth

import (
	"sentinel/internal/models"

	"github.com/google/uuid"
)

// Repository defines the base CRUD operations for any entity type T.
// This generic interface provides common database operations that should be
// available for all entities in the system.
//
// Usage example:
//
//	type PostRepository interface {
//	    Repository[models.Post]
//	    GetBySlug(slug string) (*models.Post, error)
//	}
type Repository[T any] interface {
	Create(entity *T) error
	GetByID(id uuid.UUID) (*T, error)
	Update(entity *T) error
	Delete(id uuid.UUID) error
}

// UserRepository extends the generic Repository interface with User-specific operations.
// This pattern allows for:
// 1. Common CRUD operations through Repository[models.User] embedding
// 2. Entity-specific methods like GetByEmail, Exists, etc.
// 3. Multiple storage implementations (PostgreSQL, Redis, MongoDB, etc.)
//
// Implementation example:
//   type PostgresUserRepository struct { db *database.DB }
//   func (r *PostgresUserRepository) Create(user *models.User) error { /* SQL implementation */ }
//   func (r *PostgresUserRepository) GetByEmail(email string) (*models.User, error) { /* SQL implementation */ }

type UserRepository interface {
	Repository[models.User]
	GetByEmail(email string) (*models.User, error)
	Exists(email string) (bool, error)
	UpdateLastLogin(userID uuid.UUID) error
}
