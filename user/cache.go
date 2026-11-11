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
	ms.users[u.UserName] = u
	log.Println("Create user : ", u.UserName)
	return nil
}

func (ms *MemoryStore) FindByUsername(username string) (user User, err error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	user, ok := ms.users[username]
	if !ok {
		return user, ErrNotFound
	}
	return user, nil
}
