package user

import "sync"

type MemoryStore struct {
	mu    sync.Mutex
	users map[string]User
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		users: make(map[string]User),
	}
}
