package models

import (
	"database/sql"
	"fmt"
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
	TokenManager TokenManager
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	token, tokenHash, err := ss.TokenManager.New()
	if err != nil {
		return nil, fmt.Errorf("error creating token: %w", err)
	}
	session := Session{
		UserID:    userID,
		Token:     token,
		TokenHash: tokenHash,
	}

	row := ss.DB.QueryRow(`
		UPDATE sessions
		SET token_hash = $2
		WHERE user_id = $1
		RETURNING id;
		`, session.UserID, session.TokenHash)

	err = row.Scan(&session.ID)
	if err == sql.ErrNoRows {
		// if no rows found, create a new session
		row = ss.DB.QueryRow(
			`INSERT INTO sessions (user_id, token_hash)
		 VALUES ($1, $2)
		 RETURNING id;`, session.UserID, session.TokenHash)

		err = row.Scan(&session.ID)
	}

	if err != nil {
		return nil, fmt.Errorf("saving session: %w", err)
	}

	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	tokenHash := ss.TokenManager.hash(token)

	var user User
	row := ss.DB.QueryRow(`
		SELECT email FROM users u 
		JOIN sessions s
		ON u.id = s.user_id
		WHERE s.token_hash = $1;
		`, tokenHash)
	
	err := row.Scan(&user.Email)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}
	
	return &user, nil
}

func (ss *SessionService) Delete(token string) error {
	tokenHash := ss.TokenManager.hash(token)

	// delete session from DB
	_, err := ss.DB.Exec(`
		DELETE FROM sessions
		WHERE token_hash = $1;
		`, tokenHash)

	if err != nil {
		return fmt.Errorf("error deleting session: %w", err)
	}

	return nil
}
