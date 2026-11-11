package user

import (
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
	return User{
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
