package user

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserName     string `json:"username"`
	Email        string `json:"email"`
	HashPassword string `json:"hash_password"`
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
