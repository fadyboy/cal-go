package models

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/fadyboy/lenslocked/rand"
)

const (
	// The minimum number of bytes to be used for each session token
	MinBytesPerToken = 32
)

type TokenManager struct {
	Token         string
	TokenHash     string
	BytesPerToken int
}

func (tm TokenManager) New() (token, tokenHash string, err error) {
	bytesPerToken := tm.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}

	// create token
	token, err = rand.String(bytesPerToken)
	if err != nil {
		return "", "", fmt.Errorf("error creating token: %w", err)
	}

	tokenHash = tm.hash(token)

	return token, tokenHash, nil
}

func (tm TokenManager) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.RawURLEncoding.EncodeToString(tokenHash[:])
}
