package user

import (
	"errors"
	"log"
)

type Store interface {
	Create(User) (id int, err error)
	FindByUsername(username string) (user User, err error)
}

var (
	ErrNotFound = errors.New("user not found")
	ErrExists   = errors.New("user already exists")
)

type TestStore struct {
}

func (s *TestStore) Create(user User) (id int, err error) {

	log.Printf("Create user: %s", user.UserName)
	return 0, nil
}
func (s *TestStore) FindByUsername(username string) (user User, err error) {
	return User{}, nil
}
