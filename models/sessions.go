package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"

	"github.com/fadyboy/lenslocked/rand"
)

const (
	// The minimum number of bytes to be used for each session token
	MinBytesPerToken = 32
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
	// BytesPerToken is used to determine the number of bytes when generating the session token 
	// If the value is not set or less than, it will use the MinBytesPerToken
	BytesPerToken int
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create error: %w", err)
	}

	session := Session{
		UserID: userID,
		Token: token,
		TokenHash: ss.hash(token),
	}

	row := ss.DB.QueryRow(
		`INSERT INTO sessions (user_id, token_hash)
		 VALUES ($1, $2)
		 RETURNING id;`, session.ID, session.TokenHash)

	err = row.Scan(&session.ID)
	if err != nil {
		return nil, fmt.Errorf("saving session: %w", err)
	}

	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	return nil, nil
}

func (ss *SessionService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}
