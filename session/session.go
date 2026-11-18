package session

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"
)

type Session struct {
	ID        string
	UserID    string
	CreatedAt time.Time
	ExpiresAt time.Time
}

type Store interface {
	Create(session Session) error
	Get(ID string) (Session, error)
	Delete(ID string) error
}

var ErrNotFound = errors.New("session not found")

func NewSession(userID string, expiresAt time.Duration) (Session, error) {

	return Session{
		ID:        base64.RawURLEncoding.EncodeToString([]byte(rand.Text())),
		UserID:    userID,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(expiresAt),
	}, nil
}
