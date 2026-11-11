package user

import (
	"crypto/rand"
	"encoding/hex"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string
	UserName     string `json:"username"`
	Email        string `json:"email"`
	HashPassword string `json:"hash_password"`
}

var (
	ErrNotFound = errors.New("user not found")
	ErrExists   = errors.New("user already exists")
)

type Store interface {
	Create(User) (err error)
	FindByUsername(username string) (user User, err error)
}

func NewUser(userName, email, Password string) User {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(Password), 10)
	if err != nil {
		panic(err)
	}
	id, err := generateSecureString()
	if err != nil {
		panic(err)
	}
	return User{
		ID:           id,
		UserName:     userName,
		Email:        email,
		HashPassword: string(hashPassword),
	}
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func generateSecureString() (string, error) {
	// Create a byte slice of 16 bytes
	bytes := make([]byte, 16)

	// Fill the slice with cryptographically secure random bytes
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	// Convert the 16 bytes to a 32-character hex string
	return hex.EncodeToString(bytes), nil
}
