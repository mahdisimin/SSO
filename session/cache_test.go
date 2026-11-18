package session

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestMemoryStore_Create(t *testing.T) {
	t.Helper()

	MemoryIns := NewMemoryStore()
	s, errNewSession := NewSession("user-1", time.Minute*5)
	if errNewSession != nil {
		t.Error(errNewSession)
	}
	if err := MemoryIns.Create(s); err != nil {
		t.Fatal(err)
	}

	sessionIns, errGetSession := MemoryIns.Get(s.ID)
	if errGetSession != nil {
		t.Fatal(errGetSession)
	}
	if sessionIns.UserID != "user-1" {
		t.Fatalf("invalid user id %s ", sessionIns.UserID)
	}
}

func TestMemoryStore_Delete(t *testing.T) {
	t.Helper()
	MemoryIns := NewMemoryStore()
	s, errNewSession := NewSession("user-1", time.Minute*5)
	if errNewSession != nil {
		t.Error(errNewSession)
	}
	if err := MemoryIns.Create(s); err != nil {
		t.Fatal(err)
	}
	if err := MemoryIns.Delete(s.ID); err != nil {
		t.Fatal(err)
	}
	_, errGetSession := MemoryIns.Get(s.ID)
	if errGetSession == nil {
		t.Error("expected error, got none")
	}
}

func TestMemoryStore_Expired(t *testing.T) {
	memory := NewMemoryStore()
	Newsession, err := NewSession("user-1", -1*time.Minute*5)
	if err != nil {
		t.Fatal(err)
	}
	if err := memory.Create(Newsession); err != nil {
		t.Fatal(err)
	}
	_, errGet := memory.Get(Newsession.ID)
	fmt.Println(errGet)
	if !errors.Is(errGet, ErrNotFound) {
		t.Error("expected error, got none or else")
	}
}
