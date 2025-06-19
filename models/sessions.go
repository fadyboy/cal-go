package models

import (
	"database/sql"
)

type Session struct {
	ID     int
	UserID int
	// Token is only set when creating a new Session
	// This will be left empty when looking up a session as only a hash
	// is stored in the DB and cannot be reversed into a raw token
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	return nil, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	return nil, nil
}
