package user

import (
	"log"
	"sync"
)

type MemoryStore struct {
	mu    sync.RWMutex
	users map[string]User
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		users: make(map[string]User),
	}
}

func (ms *MemoryStore) Create(u User) (err error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if _, ok := ms.users[u.UserName]; ok {
		return ErrExists
	}

	ms.users[u.UserName] = u
	log.Println("Create user : ", u.UserName)
	return nil
}

func (ms *MemoryStore) FindByUsername(username string) (user User, err error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	username = normalizeUsername(username)
	user, ok := ms.users[username]
	if !ok {
		return user, ErrNotFound
	}
	return user, nil
}

func (ms *MemoryStore) FindByID(id string) (User, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	for _, user := range ms.users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, ErrNotFound
}
