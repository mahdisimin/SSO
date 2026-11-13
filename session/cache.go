package session

import "sync"

type MemoryStore struct {
	mu    sync.RWMutex
	users map[string]Session
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{users: make(map[string]Session)}
}

func (ms *MemoryStore) Create(s Session) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.users[s.ID] = s
	return nil
}

func (ms *MemoryStore) Get(userID string) (Session, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	session, ok := ms.users[userID]
	if !ok {
		return session, ErrNotFound
	}
	return session, nil
}

func (ms *MemoryStore) Delete(userId string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	delete(ms.users, userId)
	return nil
}
