package models

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Email        string
	PasswordHash string
}

type UserService struct {
	DB *sql.DB
}

func (us *UserService) Create(email, password string) (*User, error) {
	email = strings.ToLower(email)
	passwordHashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}
	passwordHash := string(passwordHashBytes)

	user := &User{
		Email:        email,
		PasswordHash: passwordHash,
	}

	row := us.DB.QueryRow(`
		INSERT INTO users (email, password_hash)
		VALUES($1, $2) RETURNING id
		`, email, passwordHash)

	err = row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving created user from DB:%w", err)
	}

	return user, nil
}

func (us *UserService) Authenticate(email, password string) (*User, error) {
	normalizedEmail := strings.ToLower(email)
	user := &User{
		Email: normalizedEmail,
	}

	row := us.DB.QueryRow(`
		SELECT id, password_hash FROM users
		WHERE email=$1
		`, user.Email)

	err := row.Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("authentication error: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("authenticate error: %w", err)
	}

	return user, nil
}
