package auth

import (
	"database/sql"
	"fmt"
	"time"

	"sentinel/internal/models"
	"sentinel/pkg/database"

	"github.com/google/uuid"
)

// PostgresUserRepository is a PostgreSQL implementation of the UserRepository interface.
// It provides concrete implementations for both the generic Repository[models.User] methods
// and the User-specific methods defined in UserRepository.
//
// To add a new storage backend, create a new struct that implements UserRepository:
//   type RedisUserRepository struct { client *redis.Client }
//   func NewRedisUserRepository(client *redis.Client) UserRepository { return &RedisUserRepository{client} }
//   func (r *RedisUserRepository) Create(user *models.User) error { /* Redis implementation */ }
//   // ... implement all UserRepository methods
type PostgresUserRepository struct {
	db *database.DB
}

// NewUserRepository creates a new PostgreSQL-backed UserRepository.
// This factory function returns the interface type, allowing easy swapping of implementations.
func NewUserRepository(db *database.DB) UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(user *models.User) error {
	query := `
		INSERT INTO users (id, email, password_hash, first_name, last_name, is_active, is_verified, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.Exec(query,
		user.ID,
		user.Email,
		user.PasswordHash,
		user.FirstName,
		user.LastName,
		user.IsActive,
		user.IsVerified,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (r *PostgresUserRepository) GetByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, email, password_hash, first_name, last_name, is_active, is_verified, created_at, updated_at, last_login
		FROM users
		WHERE email = $1
	`

	user := &models.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.IsActive,
		&user.IsVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.LastLogin,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	return user, nil
}

func (r *PostgresUserRepository) GetByID(id uuid.UUID) (*models.User, error) {
	query := `
		SELECT id, email, password_hash, first_name, last_name, is_active, is_verified, created_at, updated_at, last_login
		FROM users
		WHERE id = $1
	`

	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.IsActive,
		&user.IsVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.LastLogin,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return user, nil
}

func (r *PostgresUserRepository) Exists(email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	
	var exists bool
	err := r.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check if user exists: %w", err)
	}
	
	return exists, nil
}

func (r *PostgresUserRepository) UpdateLastLogin(userID uuid.UUID) error {
	query := `UPDATE users SET last_login = $1, updated_at = $2 WHERE id = $3`
	
	now := time.Now()
	_, err := r.db.Exec(query, now, now, userID)
	if err != nil {
		return fmt.Errorf("failed to update last login: %w", err)
	}
	
	return nil
}

func (r *PostgresUserRepository) Update(user *models.User) error {
	query := `
		UPDATE users 
		SET email = $2, password_hash = $3, first_name = $4, last_name = $5, 
		    is_active = $6, is_verified = $7, updated_at = $8, last_login = $9
		WHERE id = $1
	`
	
	user.UpdatedAt = time.Now()
	_, err := r.db.Exec(query,
		user.ID,
		user.Email,
		user.PasswordHash,
		user.FirstName,
		user.LastName,
		user.IsActive,
		user.IsVerified,
		user.UpdatedAt,
		user.LastLogin,
	)
	
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	
	return nil
}

func (r *PostgresUserRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	
	return nil
}