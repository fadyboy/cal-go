package models

import (
	"database/sql"
	"fmt"
	"time"
)

const (
	// DefaultResetDuration is the default time that a PasswordReset is valid for
	DefaultResetDuration = 1 * time.Hour
)

type PasswordReset struct {
	ID     int
	UserID int
	// Token is only set when password reset is created
	Token     string
	TokenHash string
	ExpiresAt time.Time
}

type PasswordResetService struct {
	DB *sql.DB
	// BytesPerToken is used to determine how many bytes to user when generating
	// each reset token. If not set, we use the MinBytesPerToken constant
	BytesPerToken int
	// Duration is the amount of time a PasswordReset is valid for.
	// Defaults to DefaultResetDuration
	Duration time.Duration
}

func (prs *PasswordResetService) Create(email string) (*PasswordReset, error) {
	// stub for creating PasswordReset token
	return nil, fmt.Errorf("TODO: implement PasswordReset.Create")
}

func (prs *PasswordResetService) Consume(token string) (*User, error) {
	return nil, fmt.Errorf("TODO: implement PasswordReset.Consume")
}
